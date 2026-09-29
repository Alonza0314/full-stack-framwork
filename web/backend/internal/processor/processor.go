package processor

import (
	"backend/internal/context"
	"backend/logger"
	"time"
)

type ProcessorIE struct {
	Username string
	Password string

	JwtSecret    string
	JwtExpiresIn time.Duration

	*context.SystemContext

	*logger.BackendLogger
}

type Processor struct {
	username string
	password string

	jwtSecret    string
	jwtExpiresIn time.Duration

	*context.SystemContext

	*logger.BackendLogger
}

func NewProcessor(ie *ProcessorIE) *Processor {
	return &Processor{
		username: ie.Username,
		password: ie.Password,

		jwtSecret:    ie.JwtSecret,
		jwtExpiresIn: ie.JwtExpiresIn,

		SystemContext: ie.SystemContext,

		BackendLogger: ie.BackendLogger,
	}
}

func (p *Processor) Release() {
	p.ProcLog.Infoln("Release processor...")

	p.SystemContext.Release()

	p.ProcLog.Infoln("Processor released")
}
