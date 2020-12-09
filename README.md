callisto
========

Yet another Solar System simulator, written by Valerian Saliou in Go.

**Disclaimer: the purpose of this simulator isnt to be realistic. It was a nice way to teach myself OpenGL basics with complex 3D objects (spheres), and simple transforms. The code is shared as an example basis for those learning OpenGL; as I did.**

![Callisto - Yet another Solar System simulator](https://valeriansaliou.github.io/callisto/images/solar-system-simulator.jpg)

## Dependencies

 * **Go** (install it via: `brew install golang` on MacOS w/ Homebrew)
 * **OpenGL headers** and **GLFW headers** (built-in on MacOS)

Also, check that your `$GOPATH` is configured, and that `$GOPATH/bin` is sourced in your `$PATH`.

### MacOS

Should work out of the box.

### Linux (Ubuntu, Debian)

Install the necessary utilities and libraries:

`sudo apt-get install git libglfw-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev`

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

 * **Simulation**
  * `Scroll UP` `Scroll DOWN`: decrease/increase simulation speed

 * **Application**
  * `ESCAPE`: exit Callisto

## Disclaimer

Distances and radiuses, as well as rotation/revolution periods have been respected. Since both the Sun, Jupiter and Saturn are quite huge relative to other Solar System objects, a square-root factor has been applied on all radiuses and distances. This makes huge objects smaller on display, and small objects visible on display.

## Thanks

This project has been achieved following the excellent step-by-step tutorial available on [https://open.gl](https://open.gl)

## Copyrights

Assets (planets, moons, miscellaneous space object) are copyright NASA.
