Usage:

`scrawl testfile.tex.scr > testfile.tex`

`scrawl < testfile.tex.scr > testfile.tex`

The syntax is as follows:
Lines beginning with the triad `%-#` are called "code lines", and are interpreted by `scrawl` - all other lines are simply re-emitted.

Code lines must come in consecutive blocks, each block terminated by a line beginning
`%-#---` - anything after the third dash is ignored.
The first line, except for the `%-#`, is run using `bash`.
The remaining lines before the terminator (again, stripped of the `%-#`), are fed as standard input to this command.
The standard output is then emitted.

Note that there's almost no error checking, so don't fuck up the rules.



See the test file for an example.
