package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"

	"github.com/chenbh/concourse/atc/db/lock"
)

// A lot of struct refer to a pipeline. This is a helper interface that should
// embedded in those interfaces that need to refer to a pipeline. Accordingly,
// implementations of those interfaces should embed "pipelineRef".
type PipelineRef interface {
	PipelineID() int
	PipelineName() string
	Pipeline() (Pipeline, bool, error)
}

type pipelineRef struct {
	pipelineID   int
	pipelineName string

	conn        Conn
	lockFactory lock.LockFactory
}

func NewPipelineRef(id int, name string, conn Conn, lockFactory lock.LockFactory) PipelineRef {
	return pipelineRef{
		pipelineID:   id,
		pipelineName: name,
		conn:         conn,
		lockFactory:  lockFactory,
	}
}

func (r pipelineRef) PipelineID() int {
	return r.pipelineID
}

func (r pipelineRef) PipelineName() string {
	return r.pipelineName
}

func (r pipelineRef) Pipeline() (Pipeline, bool, error) {
	if r.PipelineID() == 0 {
		return nil, false, nil
	}

	row := pipelinesQuery.
		Where(sq.Eq{"p.id": r.PipelineID()}).
		RunWith(r.conn).
		QueryRow()

	pipeline := newPipeline(r.conn, r.lockFactory)
	err := scanPipeline(pipeline, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	return pipeline, true, nil
}
