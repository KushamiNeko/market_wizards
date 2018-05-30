//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawMarketCapitalization);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawMarketCapitalization() {

  "use strict";

  var raw = $("#market-capitalization").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Market Capitalization",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("market-capitalization"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawUpDownVolumeRatio);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawUpDownVolumeRatio() {

  "use strict";

  var raw = $("#up-down-volume-ratio").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Up Down Volume Ratio",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("up-down-volume-ratio"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
