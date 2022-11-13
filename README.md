# winterhill-lib
library to interface with the [WinterHill](https://wiki.batc.org.uk/WinterHill_Receiver_Project) board.

This software is intended to run on the pi that is part of this board.
It will connect to the local software over the exposed ports (`9901`, `9902` and `9921-9924`).
The library offers (de)serialisation of the various command and status events.

A sample application is provided which lets you control the board using a webpage.
This can be accessed locally as well as remote. For the code for this, see [app](app/) and [web](web/).


## Todo
* configure upper limit for srate per receiver in backend
* delay reset with every set
* fix layout for srate:
  * `<value> <select> <button>`; make select wider
  * only make button red if we did ask to set something, and it's not yet done. 
    Now it's also red when the receiver auto-shifts it a bit
