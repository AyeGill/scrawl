Usage:

`scrawl testfile.tex.scr > testfile.tex`

`scrawl < testfile.tex.scr > testfile.tex`

`scrawl -c testcmds testfile.tex.scr > testfile.tex`

`scrawl -v testfile.tex.scr > testfile.tex`

The syntax is as follows:
Lines beginning with the triad `%-#` are called "code lines", and are interpreted by `scrawl` - all other lines are simply re-emitted.

Code lines must come in consecutive blocks, each block terminated by a line beginning
`%-#---` - anything after the third dash is ignored.
The first line, except for the `%-#`, is run using `bash`.
The remaining lines before the terminator (again, stripped of the `%-#`), are fed as standard input to this command.
The standard output is then emitted.

Note that there's almost no error checking, so don't fuck up the rules.

The `-v` flag provides some minor debugging output.

You can provide a macro/command file with the `-c` flag.
Each line on the command file must contain a command name with no spaces, followed by a space, followed by a command (which may contain spaces, but no newlines).
A block starting with a name from the command file is run using the matching command.

See the test file for an example.
The test file can be built with `scrawl -c testcmds testfile.tex.scr > testfile.tex`.
Note that this requires the [strid](https://github.com/smimram/strid) utility.
