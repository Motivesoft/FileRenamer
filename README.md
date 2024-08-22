# FileRenamer
Simple bulk file renamer for one specific purpose but could be customised/generalised

Created for large collections of files that need their named standardised.

## Current status
* Basic framework, written in Go and therefore somewhat cross-platform
* Expects a directory path on the command line
* Iterates each file in that directory (not subdirectories)
* Applies the currently hardcoded rename rules to each file

## Current rules
* Equalising case of identifying attributes in filenames
  * e.g. "s01e02", "S01e02" and "s01E02" all become "S01E02"
* Spaces changed to dots
  * e.g. "Book and Author Name.txt" becomes "Book.and.Author.Name.txt"

## Plans
* Nothing, maybe it isn't needed
* Modifying it manually each time a specific need comes up
* Incrementally improve it with more hardcoded rules
* Add dynamic configuration capabilities
  * e.g. regex
  * e.g. 
* If we do anything at all, add:
  * version number
  * help
  * dry run mode
  * silent and verbose modes
  * published releases