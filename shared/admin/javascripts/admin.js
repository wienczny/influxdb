var recentQueries = [];
var queryPointer = null;

var data = {
    results: [{
        series: [{
            name: "cpu",
            tags: { host: "server01", region: "us-west" },
            columns: ["time", "value"],
            values: [
                ["2015-01-29T21:51:28.968422294Z", 0.64],
                ["2015-01-29T21:51:28.968422294Z", 0.64]
            ]
        }]
    }]
}

var getSeriesFromJSON = function(data) {
    var results = data.results[0].series;
    return results;
}

var handleSubmit = function(e) {
    var queryElement = document.getElementById('query');
    var q = queryElement.value;

    var query = $.get("http://localhost:8086/query", {q: q, db: "mydb"}, function() {
        console.log("query sent!");
    });

    recentQueries.push(q);
    queryPointer = recentQueries.length - 1;

    query.fail(function (e) {
        if (e.status == 400) {
            var response = JSON.parse(e.responseText)
            React.render(
              React.createElement(QueryError, {message: response.error}),
              document.getElementById('alert')
            );
        }
    });

    query.done(function (data) {
        var firstRow = data.results[0];
        if (firstRow.error) {
            React.render(
              React.createElement(QueryError, {message: firstRow.error}),
              document.getElementById('alert')
            );
            return
        }

        var series = getSeriesFromJSON(data);

        React.render(
          React.createElement(DataTable, {series: series}),
          document.getElementById('table')
        );
    });

    e.preventDefault();
    return false;
};

var handleKeypress = function(e) {
    var queryElement = document.getElementById('query');
    if (recentQueries.length == 0 ) { return }

    if (e.keyCode == 38) {
        // stash the current query, if there is one?

        if (queryPointer != null && queryPointer > 0) {
            queryPointer -= 1;
            queryElement.value = recentQueries[queryPointer];
        }
    }

    if (e.keyCode == 40) {
        if (queryPointer != null && queryPointer < recentQueries.length - 1) {
            queryPointer += 1;
            queryElement.value = recentQueries[queryPointer];
        }
    }
};

var QueryError = React.createClass({
    render: function() {
        return React.createElement("div", {className: "alert alert-danger"}, this.props.message)
    }
});

var QueryField = React.createClass({
  render: function() {
    return (
    return React.createElement("form", {id: "query-form"}, 
        React.createElement("div", {className: "form-group"}, 
          React.createElement("input", {type: "text", className: "form-control", id: "query"})
        )
      )
            
      <form id="query-form">
        <div className="form-group">
          <input type="text" className="form-control" id="query" />
        </div>
      </form>
    )
  },

  componentDidMount: function() {
    var form = document.getElementById('query-form');
    form.addEventListener("submit", handleSubmit);

    var query = document.getElementById('query');
    query.addEventListener("keydown", handleKeypress);

  }
});

var DataTable = React.createClass({
  render: function() {
    console.log(this.props.series)

    var tables = this.props.series.map(function(series) {
        return React.createElement("div", null, 
            React.createElement("h1", null, series.name), 
            React.createElement("table", {className: "table"}, 
                React.createElement(TableHeader, {data: series.columns}), 
                React.createElement(TableBody, {data: series})
            )
        );
    });

    return React.createElement("div", null, tables);
  }
});

var TableHeader = React.createClass({
    render: function() {
        var headers = this.props.data.map(function(column) {
            return React.createElement("th", null, column);
        });

        return React.createElement("tr", null, headers);
    }
});

var TableBody = React.createClass({
    render: function() {
        var tableRows = this.props.data.values.map(function (row) {
            return React.createElement(TableRow, {data: row});
        });

        return React.createElement("tbody", null, tableRows);
    }
});

var TableRow = React.createClass({
    render: function() {
        console.log("data:", data)
        var tableData = this.props.data.map(function (data) {
            return React.createElement("td", null, data)
        });

        return React.createElement("tr", null, tableData);
    }
});

React.render(
  React.createElement(QueryField, null),
  document.getElementById('content')
);

document.getElementById('query').focus();

