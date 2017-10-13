# Change Log


## Unreleased

### Added

- Daemon and cron tools
- `WithTimeout` function to enforce timeouts
- Simple http server construction
- Database connection constructor
- Master-slave database connection constructor
- Add gRPC server constructor

### Changed

- Do not require a `net.Addr` in debug server config
- Debug server is started without goph/serverz
- Make configurations non-optional options


## 0.2.0 - 2017-10-12

### Added

- Error channel returned from debug server
- Opentracing global tracer

### Changed

- Provide debug.Handler from debug server constructor
- Make logger optional (but highly recommended) for debug server

### Removed

- Global error channel
- Debug bootstrap
- goph/stdlib requirement


## 0.1.0 - 2017-10-12

### Added

- [go-kit](https://github.com/go-kit/kit/tree/master/log) logger constructor
- [emperror](https://github.com/goph/emperror) error handler constructor
- [airbrake](https://github.com/airbrake/gobrake) compatible error handler constructor
- Error handler stack for collecting multiple handlers under a single dependency
- [opentracing](http://opentracing.io/) tracer constructor
- [jaeger](https://github.com/jaegertracing) tracer constructor
- Lifecycle extension for supporting Closers
- Debug server constructor (invoke fn)
- Health collector constructor
- Bootstrap options
