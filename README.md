callisto
========

[![Build Status](https://img.shields.io/travis/valeriansaliou/callisto/master.svg)](https://travis-ci.org/valeriansaliou/callisto)

Yet another Solar System simulator, written by Valerian Saliou in Go.

**Disclaimer: the purpose of this simulator isnt to be realistic. It was a nice way to teach myself OpenGL basics with complex 3D objects (spheres), and simple transforms. The code is shared as an example basis for those learning OpenGL; as I did.**

![Callisto - Yet another Solar System simulator](https://valeriansaliou.github.io/callisto/images/solar-system-simulator.jpg)

## Dependencies

 * **Go** (install it via: `brew install golang` on MacOS w/ Homebrew)
 * **OpenGL headers** and **GLFW headers** (built-in on MacOS)

Also, check that your `$GOPATH` is configured, and that `$GOPATH/bin` is sourced in your `$PATH`.

## Install & Run

 1. `go get github.com/valeriansaliou/callisto`
 2. `callisto`

## Controls

 * **Direction**
  * `UP` `DOWN` `LEFT` `RIGHT`: move camera position
  * `Mouse move`: move camera target

 * **Camera**
  * `R`: reset camera position
  * `SPACE`: turbo camera move (hold key)
  * `TAB`: toggle camera view (from flight view to stellar object view lock)

 * **Simulation**
  * `Scroll UP` `Scroll DOWN`: decrease/increase simulation speed

## Disclaimer

Distances and radiuses, as well as rotation/revolution periods have been respected. Since both the Sun, Jupiter and Saturn are quite huge relative to other Solar System objects, a square-root factor has been applied on all radiuses and distances. This makes huge objects smaller on display, and small objects visible on display.

The main asteroid belt does not reflect reality, it has been trimmed down to some random objects only, while in reality it should contain at least 7,000 known asteroids. The Kuiper asteroid belt has been omitted for performance reasons (it should contain at least a trillon objects, which I doubt any GPU can handle in 2016).

## Thanks

This project has been achieved following the excellent step-by-step tutorial available on [https://open.gl](https://open.gl)

## Copyrights

Assets (planets, moons, miscellaneous space object) are copyright NASA.
