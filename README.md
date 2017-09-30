# pacchetto

## What is it?
pacchetto is a tool that creates a minimal archive file containing the files
required to run Assetto Corsa dedicated server on another machine. This is
accomplished in two layers:

* A Go library
* A Go command line application

## What does that mean?
That you can create an archive and upload it to your Assetto Corsa host.
This means that you do not have to download and sign into Steam on the host.
It also allows you to store clean copies of the game data somewhere.

## How do I use it?
1. Install Assetto Corsa on a "staging" computer
2. Either build or download the pacchetto application on the staging machine
3. Open the Command Prompt (Windows) or your favorite terminal application
4. `cd` to the directory containing your pacchetto executable
5. Run `pacchetto -p` and wait for the package creation to finish

This will create an archive that can then be transferred to another machine.

## How do I build it?
Please refer to the [building documentation](docs/building/README.md).