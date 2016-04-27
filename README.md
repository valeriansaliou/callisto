callisto
========

Yet another Solar System simulator, written by Valerian Saliou in Go.

**Disclaimer: the purpose of this simulator isnt to be realistic. It was a nice way to teach myself OpenGL basics with complex 3D objects (spheres), and simple transforms. The code is shared as an example basis for those learning OpenGL; as I did.**

![Callisto - Yet another Solar System simulator](https://valeriansaliou.github.io/callisto/images/solar-system-simulator.jpg)

## Dependencies

 * **Go** (install it via: `brew install golang` on MacOS w/ Homebrew)
 * **OpenGL headers** and **GLFW headers** (built-in on MacOS)

Also, check that your `$GOPATH` is configured, and that you source `$GOPATH/bin` in your `$PATH`.

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

## Thanks

This project has been achieved following the excellent step-by-step tutorial available on [https://open.gl](https://open.gl)

## Copyrights

Assets (planets, moons, miscellaneous space object) are copyright NASA.
