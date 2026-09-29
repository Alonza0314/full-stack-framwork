package context

import "backend/logger"

type SystemContextIE struct {
	DbType string
	DbPath string

	*logger.BackendLogger
}

type SystemContext struct {
	*dbContext

	*logger.BackendLogger
}

func NewSystemContext(ie *SystemContextIE) *SystemContext {
	dbContext, err := newDbContext(&dbContextIE{
		dbType: ie.DbType,
		dbPath: ie.DbPath,

		BackendLogger: ie.BackendLogger,
	})
	if err != nil {
		ie.BackendLogger.CtxLog.Errorf("Failed to create dbContext: %v", err)
		return nil
	}

	return &SystemContext{
		dbContext: dbContext,

		BackendLogger: ie.BackendLogger,
	}
}

func (ctx *SystemContext) Release() {
	ctx.CtxLog.Infoln("Release SystemContext...")

	ctx.dbContext.release()

	ctx.CtxLog.Infoln("SystemContext released")
}
