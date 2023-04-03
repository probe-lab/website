# ashby

`ashby` produces plots for the cmi website

Currently it only supports an interactive mode but the goal is to run it as a daemon and for it to 
produce automatically plots on a schedule.

## Getting Started

Run `go build` in the `/cmd/ashby` folder.

If you're just trying the static demo plots then you can use:

	./ashby plot --preview ../../plots/demo-static-bar-grouped.json

It outputs the JSON that defines the plot and can be passed to the plotly JavaScript library.

`--preview` makes ashby launch the plot as a preview in your browser

If you want to try a postgres demo plot then you need to specify the postgres datasource with the `-s` option:

	-s name=postgres://username:password@hostname:5432/database_name

For the demos the name is `pgnebula`, so you should run:

	./ashby plot --preview -s pgnebula=postgres://username:password@hostname/database_name?sslmode=prefer ../../plots/demo-pg-agents-bar.json

(with the database details filled in, of course!)


## Plot Specifications

Plots are defined in JSON. Some samples are in the `plots` folder at the root of this repo.

The specification format is in flux but currently there are three sections to each plot specification.

 - `datasets` - this specifies a list of named datasets that provide source data for the plots. Each has a source and query. Support for query parameters is on the TODO list. A dataset is essentially a list of named fields and their data values, usually a tabular structure.
 - `series` - this specifies a list of series that are to be plotted. Each series specifies the field to use for labelling the points in the series and a field for the values. Each series will be plotted onto the final chart.
 - `layout` - this defines the layout for the plot. Currently it's just the same as the plotly layout definition but ideally we will support only a useful subset to avoid coupling too tightly to a single plotting library.


An example plot spec:

```json
	{
	  "datasets": [
	    {
	      "name":"main",
	      "source":"demo",
	      "query":"populations"
	    }
	  ],

	  "series": [
	    {
	      "type": "bar",
	      "name": "population",
	      "dataset": "main",
	      "labels": "creature",
	      "values": "month1"
	    }
	  ],

	  "layout":{
	    "title":{
	      "text":"Demo: Bar"
	    }
	  }
	}
```
