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

google.charts.setOnLoadCallback(drawRSRating);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawRSRating() {

  "use strict";

  var raw = $("#rs-rating").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "RS Rating",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("rs-rating"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawIndustryGroupRank);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawIndustryGroupRank() {

  "use strict";

  var raw = $("#industry-group-rank").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Industry Group Rank",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("industry-group-rank"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawCompositeRating);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawCompositeRating() {

  "use strict";

  var raw = $("#composite-rating").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Composite Rating",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("composite-rating"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawEPSRating);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawEPSRating() {

  "use strict";

  var raw = $("#eps-rating").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "EPS Rating",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("eps-rating"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawSMRRating);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawSMRRating() {

  "use strict";

  var raw = $("#smr-rating").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "SMR Rating",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("smr-rating"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawAccDisRating);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawAccDisRating() {

  "use strict";

  var raw = $("#acc-dis-rating").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Accumulation/Distribution Rating",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("acc-dis-rating"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
