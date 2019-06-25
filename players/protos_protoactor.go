
package players

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var rootContext = actor.EmptyRootContext
	
var xPlayerFactory func() Player

// PlayerFactory produces a Player
func PlayerFactory(factory func() Player) {
	xPlayerFactory = factory
}

// GetPlayerGrain instantiates a new PlayerGrain with given ID
func GetPlayerGrain(id string) *PlayerGrain {
	return &PlayerGrain{ID: id}
}

// Player interfaces the services available to the Player
type Player interface {
	Init(id string)
	Terminate()
		
	CompareHash(*CompareHashRequest, cluster.GrainContext) (*CompareHashResponse, error)
		
	ChangeLook(*LookRequest, cluster.GrainContext) (*LookResponse, error)
		
	ChangeName(*NameRequest, cluster.GrainContext) (*NameResponse, error)
		
	ChangePassword(*ChangePasswordRequest, cluster.GrainContext) (*PasswordResponse, error)
		
	ChangeMotto(*MottoRequest, cluster.GrainContext) (*MottoResponse, error)
		
	IncreaseCurrency(*CurrencyRequest, cluster.GrainContext) (*CurrencyResponse, error)
		
	DecreaseCurrency(*CurrencyRequest, cluster.GrainContext) (*CurrencyResponse, error)
		
}

// PlayerGrain holds the base data for the PlayerGrain
type PlayerGrain struct {
	ID string
}
	
// CompareHash requests the execution on to the cluster using default options
func (g *PlayerGrain) CompareHash(r *CompareHashRequest) (*CompareHashResponse, error) {
	return g.CompareHashWithOpts(r, cluster.DefaultGrainCallOptions())
}

// CompareHashWithOpts requests the execution on to the cluster
func (g *PlayerGrain) CompareHashWithOpts(r *CompareHashRequest, opts *cluster.GrainCallOptions) (*CompareHashResponse, error) {
	fun := func() (*CompareHashResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 0, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &CompareHashResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *CompareHashResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// CompareHashChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) CompareHashChan(r *CompareHashRequest) (<-chan *CompareHashResponse, <-chan error) {
	return g.CompareHashChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// CompareHashChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) CompareHashChanWithOpts(r *CompareHashRequest, opts *cluster.GrainCallOptions) (<-chan *CompareHashResponse, <-chan error) {
	c := make(chan *CompareHashResponse)
	e := make(chan error)
	go func() {
		res, err := g.CompareHashWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// ChangeLook requests the execution on to the cluster using default options
func (g *PlayerGrain) ChangeLook(r *LookRequest) (*LookResponse, error) {
	return g.ChangeLookWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeLookWithOpts requests the execution on to the cluster
func (g *PlayerGrain) ChangeLookWithOpts(r *LookRequest, opts *cluster.GrainCallOptions) (*LookResponse, error) {
	fun := func() (*LookResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 1, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &LookResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *LookResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// ChangeLookChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) ChangeLookChan(r *LookRequest) (<-chan *LookResponse, <-chan error) {
	return g.ChangeLookChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeLookChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) ChangeLookChanWithOpts(r *LookRequest, opts *cluster.GrainCallOptions) (<-chan *LookResponse, <-chan error) {
	c := make(chan *LookResponse)
	e := make(chan error)
	go func() {
		res, err := g.ChangeLookWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// ChangeName requests the execution on to the cluster using default options
func (g *PlayerGrain) ChangeName(r *NameRequest) (*NameResponse, error) {
	return g.ChangeNameWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeNameWithOpts requests the execution on to the cluster
func (g *PlayerGrain) ChangeNameWithOpts(r *NameRequest, opts *cluster.GrainCallOptions) (*NameResponse, error) {
	fun := func() (*NameResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 2, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &NameResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *NameResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// ChangeNameChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) ChangeNameChan(r *NameRequest) (<-chan *NameResponse, <-chan error) {
	return g.ChangeNameChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeNameChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) ChangeNameChanWithOpts(r *NameRequest, opts *cluster.GrainCallOptions) (<-chan *NameResponse, <-chan error) {
	c := make(chan *NameResponse)
	e := make(chan error)
	go func() {
		res, err := g.ChangeNameWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// ChangePassword requests the execution on to the cluster using default options
func (g *PlayerGrain) ChangePassword(r *ChangePasswordRequest) (*PasswordResponse, error) {
	return g.ChangePasswordWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangePasswordWithOpts requests the execution on to the cluster
func (g *PlayerGrain) ChangePasswordWithOpts(r *ChangePasswordRequest, opts *cluster.GrainCallOptions) (*PasswordResponse, error) {
	fun := func() (*PasswordResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 3, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &PasswordResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *PasswordResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// ChangePasswordChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) ChangePasswordChan(r *ChangePasswordRequest) (<-chan *PasswordResponse, <-chan error) {
	return g.ChangePasswordChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangePasswordChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) ChangePasswordChanWithOpts(r *ChangePasswordRequest, opts *cluster.GrainCallOptions) (<-chan *PasswordResponse, <-chan error) {
	c := make(chan *PasswordResponse)
	e := make(chan error)
	go func() {
		res, err := g.ChangePasswordWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// ChangeMotto requests the execution on to the cluster using default options
func (g *PlayerGrain) ChangeMotto(r *MottoRequest) (*MottoResponse, error) {
	return g.ChangeMottoWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeMottoWithOpts requests the execution on to the cluster
func (g *PlayerGrain) ChangeMottoWithOpts(r *MottoRequest, opts *cluster.GrainCallOptions) (*MottoResponse, error) {
	fun := func() (*MottoResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 4, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &MottoResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *MottoResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// ChangeMottoChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) ChangeMottoChan(r *MottoRequest) (<-chan *MottoResponse, <-chan error) {
	return g.ChangeMottoChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// ChangeMottoChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) ChangeMottoChanWithOpts(r *MottoRequest, opts *cluster.GrainCallOptions) (<-chan *MottoResponse, <-chan error) {
	c := make(chan *MottoResponse)
	e := make(chan error)
	go func() {
		res, err := g.ChangeMottoWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// IncreaseCurrency requests the execution on to the cluster using default options
func (g *PlayerGrain) IncreaseCurrency(r *CurrencyRequest) (*CurrencyResponse, error) {
	return g.IncreaseCurrencyWithOpts(r, cluster.DefaultGrainCallOptions())
}

// IncreaseCurrencyWithOpts requests the execution on to the cluster
func (g *PlayerGrain) IncreaseCurrencyWithOpts(r *CurrencyRequest, opts *cluster.GrainCallOptions) (*CurrencyResponse, error) {
	fun := func() (*CurrencyResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 5, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &CurrencyResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *CurrencyResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// IncreaseCurrencyChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) IncreaseCurrencyChan(r *CurrencyRequest) (<-chan *CurrencyResponse, <-chan error) {
	return g.IncreaseCurrencyChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// IncreaseCurrencyChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) IncreaseCurrencyChanWithOpts(r *CurrencyRequest, opts *cluster.GrainCallOptions) (<-chan *CurrencyResponse, <-chan error) {
	c := make(chan *CurrencyResponse)
	e := make(chan error)
	go func() {
		res, err := g.IncreaseCurrencyWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	
// DecreaseCurrency requests the execution on to the cluster using default options
func (g *PlayerGrain) DecreaseCurrency(r *CurrencyRequest) (*CurrencyResponse, error) {
	return g.DecreaseCurrencyWithOpts(r, cluster.DefaultGrainCallOptions())
}

// DecreaseCurrencyWithOpts requests the execution on to the cluster
func (g *PlayerGrain) DecreaseCurrencyWithOpts(r *CurrencyRequest, opts *cluster.GrainCallOptions) (*CurrencyResponse, error) {
	fun := func() (*CurrencyResponse, error) {
			pid, statusCode := cluster.Get(g.ID, "Player")
			if statusCode != remote.ResponseStatusCodeOK {
				return nil, fmt.Errorf("get PID failed with StatusCode: %v", statusCode)
			}
			bytes, err := proto.Marshal(r)
			if err != nil {
				return nil, err
			}
			request := &cluster.GrainRequest{MethodIndex: 6, MessageData: bytes}
			response, err := rootContext.RequestFuture(pid, request, opts.Timeout).Result()
			if err != nil {
				return nil, err
			}
			switch msg := response.(type) {
			case *cluster.GrainResponse:
				result := &CurrencyResponse{}
				err = proto.Unmarshal(msg.MessageData, result)
				if err != nil {
					return nil, err
				}
				return result, nil
			case *cluster.GrainErrorResponse:
				return nil, errors.New(msg.Err)
			default:
				return nil, errors.New("unknown response")
			}
		}
	
	var res *CurrencyResponse
	var err error
	for i := 0; i < opts.RetryCount; i++ {
		res, err = fun()
		if err == nil || err.Error() != "future: timeout" {
			return res, err
		} else if opts.RetryAction != nil {
				opts.RetryAction(i)
		}
	}
	return nil, err
}

// DecreaseCurrencyChan allows to use a channel to execute the method using default options
func (g *PlayerGrain) DecreaseCurrencyChan(r *CurrencyRequest) (<-chan *CurrencyResponse, <-chan error) {
	return g.DecreaseCurrencyChanWithOpts(r, cluster.DefaultGrainCallOptions())
}

// DecreaseCurrencyChanWithOpts allows to use a channel to execute the method
func (g *PlayerGrain) DecreaseCurrencyChanWithOpts(r *CurrencyRequest, opts *cluster.GrainCallOptions) (<-chan *CurrencyResponse, <-chan error) {
	c := make(chan *CurrencyResponse)
	e := make(chan error)
	go func() {
		res, err := g.DecreaseCurrencyWithOpts(r, opts)
		if err != nil {
			e <- err
		} else {
			c <- res
		}
		close(c)
		close(e)
	}()
	return c, e
}
	

// PlayerActor represents the actor structure
type PlayerActor struct {
	inner Player
	Timeout *time.Duration
}

// Receive ensures the lifecycle of the actor for the received message
func (a *PlayerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.inner = xPlayerFactory()
		id := ctx.Self().Id
		a.inner.Init(id[7:]) // skip "remote$"
		if a.Timeout != nil {
			ctx.SetReceiveTimeout(*a.Timeout)
		}
	case *actor.ReceiveTimeout:
		a.inner.Terminate()
		ctx.Self().Poison()

	case actor.AutoReceiveMessage: // pass
	case actor.SystemMessage: // pass

	case *cluster.GrainRequest:
		switch msg.MethodIndex {
			
		case 0:
			req := &CompareHashRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.CompareHash(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 1:
			req := &LookRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.ChangeLook(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 2:
			req := &NameRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.ChangeName(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 3:
			req := &ChangePasswordRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.ChangePassword(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 4:
			req := &MottoRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.ChangeMotto(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 5:
			req := &CurrencyRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.IncreaseCurrency(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 6:
			req := &CurrencyRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.DecreaseCurrency(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
		
		}
	default:
		log.Printf("Unknown message %v", msg)
	}
}

	



