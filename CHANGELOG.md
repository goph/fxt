# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- App config for SQL
- App config for Airbrake
- App config for Jaeger
- App config for gRPC client


## [0.6.3] - 2017-12-05

### Added

- `io.Writer` adapter for go-kit logger

### Changed

- Use `io.Writer` log adapter instead of the stdlib one in grpc log


## [0.6.2] - 2017-11-17

### Added

- Error handler result for error handler groups


## [0.6.1] - 2017-11-16

### Added

- StatusChecker health check constructor


## [0.6.0] - 2017-11-16

## Changed

- Upgrade [go.uber.org/dig](http://go.uber.org/dig) to minimum 1.2.0
- Use dig 1.2 features in error handling


## [0.5.2] - 2017-10-30

### Changed

- Updated dependencies


## [0.5.1] - 2017-10-26

### Added

- Stack trace to some errors

### Changed

- Updated dependencies


## [0.5.0] - 2017-10-18

### Added

- Prometheus metrics endpoint register function
- Custom gRPC server options to the gRPC config

### Changed

- Moved Gorilla Mux router to `http/gorilla` package
- Replaced custom router solution with tracer injection

### Removed

- Custom Mux router wrapper
- Interceptors from gRPC constructor params


## [0.4.0] - 2017-10-15

### Added

- Router wrapping Gorilla mux adding OpenTracing support

### Changed

- Logger is now optional in Jaeger tracer


## [0.3.3] - 2017-10-13

### Fixed

- Null pointer closer


## [0.3.2] - 2017-10-13

### Fixed

- Null pointer closer


## [0.3.1] - 2017-10-13

### Changed

- Lock go-kit to patch versions
- Register the daemon runner in the application lifecycle


## [0.3.0] - 2017-10-13

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


## [0.2.0] - 2017-10-12

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


[Unreleased]: https://github.com/goph/fxt/compare/v0.6.2...HEAD
[0.6.1]: https://github.com/goph/fxt/compare/v0.6.1...v0.6.2
[0.6.1]: https://github.com/goph/fxt/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/goph/fxt/compare/v0.5.2...v0.6.0
[0.5.2]: https://github.com/goph/fxt/compare/v0.5.1...v0.5.2
[0.5.1]: https://github.com/goph/fxt/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/goph/fxt/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/goph/fxt/compare/v0.3.3...v0.4.0
[0.3.3]: https://github.com/goph/fxt/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/goph/fxt/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/goph/fxt/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/goph/fxt/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/goph/fxt/compare/v0.1.0...v0.2.0
