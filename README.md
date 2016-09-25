# StackedBoxes's Game Engine in Go

Yet another game engine. Or someday, maybe, this will be one. For now, it is
just nothing.

Quick facts:

* Statistics predict it will never be completed :-)
* Intended to be a "primarily 2D" engine
    * Likely support "generic OpenGL" code, but that's not the focus.
* Goal is to be a high-level engine (but not one specific for a genre or two).
  Maybe more like a toolkit than a real engine?
    * My ideal: make it nice to prototype ideas (even crazy ones) and for games
      that aren't particularly resource intensive.
* Uses Allegro 5 as the "back end"
    * Will allow me to get things running quickly.
* Will not care about GC or CGO performance until it becomes an issue. Maybe
  then it will be too late to fix. *Sigh*.
* If you dare trying this lib, you better make all calls from the main
  goroutine.
