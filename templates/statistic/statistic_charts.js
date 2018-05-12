//////////////////////////////////////////////////////////////////////////////////////////////////////

var width = $("#button-type").width();

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.load("current", {
  "packages": ["corechart", "bar"]
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawGainVsDaysHeld);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawGainVsDaysHeld() {

  "use strict";

  var raw = $("#gain-daysheld").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var options = {
    width: width,
    height: width / 2,
    title: "Gain(%) vs DaysHeld",
    hAxis: {
      title: "DaysHeld",
    },
    vAxis: {
      title: "Gain(%)",
    },
    legend: "none"
  };

  var chart = new google.visualization.ScatterChart(document.getElementById("gain-daysheld"));

  chart.draw(data, options);
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawBuyPoints);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawBuyPoints() {

  "use strict";

  var raw = $("#buy-points").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var options = {
    width: width,
    height: width / 2,
    //chart: {
    //title: "Buy Points"
    //}
    //width: 900,
    series: {
      0: {
        targetAxisIndex: 0
      },
      1: {
        targetAxisIndex: 1
      }
    },
    title: 'Nearby galaxies - distance on the left, brightness on the right',
    vAxes: {
      // Adds titles to each axis.
      0: {
        title: 'parsecs'
      },
      1: {
        title: 'apparent magnitude'
      }
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("buy-points"));
  classicChart.draw(data, options);

  //var chart = new google.charts.Bar(document.getElementById("buy-points"));

  //chart.draw(data, google.charts.Bar.convertOptions(options));
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
