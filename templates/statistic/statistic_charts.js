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

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Buy Points",
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("buy-points"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
