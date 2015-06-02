var recentQueries = [];
var queryPointer = null;

var getSeriesFromJSON = function(data) {
    var results = data.results[0].series;
    return results;
}

var handleSubmit = function(e) {
    var queryElement = document.getElementById('query');
    var q = queryElement.value;

    var query = $.get("http://localhost:8086/query", {q: q, db: "mydb"}, function() {
        // TODO: maybe we start a spinner here?

        document.getElementById('alert').innerHTML = "";
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

    // key press up
    if (e.keyCode == 38) {
        // TODO: stash the current query, if there is one?
        if (queryPointer == recentQueries.length - 1) {
            // this is buggy.
            //recentQueries.push(queryElement.value);
            //queryPointer = recentQueries.length - 1;
        }

        if (queryPointer != null && queryPointer > 0) {
            queryPointer -= 1;
            queryElement.value = recentQueries[queryPointer];
        }
    }

    // key press down
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
    return React.createElement("form", {id: "query-form"}, 
        React.createElement("div", {className: "form-group"}, 
          React.createElement("input", {type: "text", className: "form-control", id: "query"})
        )
      )
  },

  componentDidMount: function() {
    var form = document.getElementById('query-form');
    form.addEventListener("submit", handleSubmit);

    var query = document.getElementById('query');
    query.addEventListener("keydown", handleKeypress);

    document.getElementById('query').focus();
  }
});

var DataTable = React.createClass({
  render: function() {
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

// TODO: make a deleteable option?
var TableRow = React.createClass({
    render: function() {
        var tableData = this.props.data.map(function (data) {
            return React.createElement("td", null, data)
        });

        return React.createElement("tr", null, tableData);
    }
});

// TODO: wrap this in a document.ready?
React.render(
  React.createElement(QueryField, null),
  document.getElementById('content')
);


