# Termgraph: A golang terminal control library
Termgraph is used to create simple terminal interfaces and perform terminal
control while focusing on both simplicity and functionality.\

This uses the idea of a Screen and Areas upon a screen to print individual
characters and whole strings, with ownership mechanics on each "cell" of the
screen, which imposes rules on what Area can be used to print. This allows 
for a simple interface to define how strings should appear on the screen

# Installation
Use the following command to install the package\
```$ go install github.com:NatePruitt1/termgraph.git```
