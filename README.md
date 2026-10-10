# tinygo-pullup

Pull-up resistor wrapper for TinyGo projects.

**Note:** This repository is archived and read-only.

## Features

- **`Handler`** — Interface with `Setup`, `IsHigh`, `IsLow`, `IsOpen` and `IsShorted`; `NewDefaultHandler(inputPin)` builds the default implementation.
- **`enabler`** — Subpackage with `IsEnabled` / `IsDisabled` handlers (`NewDefaultHandler`, `NewDefaultHandlerFromPin`).

## License

GPL-3.0. See [LICENSE](LICENSE).
