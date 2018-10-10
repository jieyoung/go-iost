package vm

import (
	"time"

	"sync"

	"fmt"

	"github.com/iost-official/go-iost/core/block"
	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/vm/database"
)

type Batcher interface {
	Batch(bh *block.BlockHead, db database.IMultiValue, provider Provider, limit time.Duration, thread int) *Batch
	Verify(bh *block.BlockHead, db database.IMultiValue, checkFunc func(e Engine, t *tx.Tx, r *tx.TxReceipt) error, b *Batch) error
}

type Batch struct {
	Txs      []*tx.Tx
	Receipts []*tx.TxReceipt
}

func NewBatch() *Batch {
	return &Batch{
		Txs:      make([]*tx.Tx, 0),
		Receipts: make([]*tx.TxReceipt, 0),
	}
}

type Provider interface {
	Tx() *tx.Tx
	Return(*tx.Tx)
}

type batcherImpl struct {
	wait sync.WaitGroup
}

func NewBatcher() Batcher {
	return &batcherImpl{}
}

func (m *batcherImpl) Batch(bh *block.BlockHead, db database.IMultiValue, provider Provider, limit time.Duration, thread int) *Batch {
	var (
		mappers  = make([]map[string]database.Access, thread)
		txs      = make([]*tx.Tx, thread)
		receipts = make([]*tx.TxReceipt, thread)
		visitors = make([]*database.Visitor, thread)
	)

	bvr := database.NewBatchVisitorRoot(10000, db)
	for i := 0; i < thread; i++ {
		i2 := i
		go func() {
			vi, mapper := database.NewBatchVisitor(bvr)

			e := newEngine(bh, vi)

			m.wait.Add(1)
			defer m.wait.Done()

			// todo setup engine=
			t := provider.Tx()
			tr, err := e.Exec(t, limit)

			if err == nil {
				mappers[i2] = mapper.Map()
				txs[i2] = t
				receipts[i2] = tr
				visitors[i2] = vi
			} else {
				mappers[i2] = nil
			}
		}()
	}
	m.wait.Wait()
	ti, td := Resolve(mappers)

	for _, i := range td {
		provider.Return(txs[i])
	}

	b := NewBatch()

	for _, i := range ti {
		b.Txs = append(b.Txs, txs[i])
		b.Receipts = append(b.Receipts, receipts[i])
		visitors[i].Commit()
	}

	return b
}

func Resolve(mappers []map[string]database.Access) (accept, drop []int) {
	workMap := make(map[string]database.Access)
	accept = make([]int, 0)
	drop = make([]int, 0)
L:
	for i, m := range mappers {
		if m == nil {
			continue
		}
		for k, v := range m {
			x, ok := workMap[k]
			switch {
			case !ok:
				workMap[k] = v
			case x == database.Read && v == database.Read:
			case x == database.Write || v == database.Write:
				drop = append(drop, i)
				continue L
			}
		}
		accept = append(accept, i)
	}
	return
}

func (v *batcherImpl) Verify(bh *block.BlockHead, db database.IMultiValue, checkFunc func(e Engine, t *tx.Tx, r *tx.TxReceipt) error, b *Batch) error {
	var (
		thread  = len(b.Txs)
		mappers = make([]map[string]database.Access, thread)
		errs    = make([]error, thread)
	)

	bvr := database.NewBatchVisitorRoot(10000, db)
	for i := 0; i < thread; i++ {
		i2 := i
		go func() {
			vi, mapper := database.NewBatchVisitor(bvr)

			e := newEngine(bh, vi)

			v.wait.Add(1)
			defer v.wait.Done()

			// todo setup engine=
			t := b.Txs[i2]
			errs[i2] = checkFunc(e, t, b.Receipts[i2])
			if errs[i2] == nil {
				mappers[i2] = mapper.Map()
			}

		}()
	}
	v.wait.Wait()
	_, td := Resolve(mappers)
	if len(td) != 0 {
		return fmt.Errorf("transaction conflicted")
	}
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}
