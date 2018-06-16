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

google.charts.setOnLoadCallback(drawEPSChgLastQtr);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawEPSChgLastQtr() {

  "use strict";

  var raw = $("#eps-chg-last-qtr").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "EPS % Chg (Last Qtr)",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("eps-chg-last-qtr"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawLast3QtrsAvgEPSGrowth);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawLast3QtrsAvgEPSGrowth() {

  "use strict";

  var raw = $("#last-3-qtrs-avg-eps-growth").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Last 3 Qtrs Avg EPS Growth",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("last-3-qtrs-avg-eps-growth"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawQtrsofEPSAcceleration);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawQtrsofEPSAcceleration() {

  "use strict";

  var raw = $("#qtrs-of-eps-acceleration").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "# Qtrs of EPS Acceleration",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("qtrs-of-eps-acceleration"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawEPSEstChgCurrentQtr);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawEPSEstChgCurrentQtr() {

  "use strict";

  var raw = $("#eps-est-chg-current-qtr").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "EPS Est % Chg (Current Qtr)",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("eps-est-chg-current-qtr"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawEstimateRevisions);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawEstimateRevisions() {

  "use strict";

  var raw = $("#estimate-revisions").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Estimate Revisions",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("estimate-revisions"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawLastQtrEarningsSurprise);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawLastQtrEarningsSurprise() {

  "use strict";

  var raw = $("#last-qtr-earnings-surprise").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Last Quarter % Earnings Surprise",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("last-qtr-earnings-surprise"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawThreeYrEPSGrowthRate);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawThreeYrEPSGrowthRate() {

  "use strict";

  var raw = $("#three-yr-eps-growth-rate").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "3 Yr EPS Growth Rate",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("three-yr-eps-growth-rate"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawConsecutiveYrsofAnnualEPSGrowth);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawConsecutiveYrsofAnnualEPSGrowth() {

  "use strict";

  var raw = $("#consecutive-yrs-of-annual-eps-growth").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Consecutive Yrs of Annual EPS Growth",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("consecutive-yrs-of-annual-eps-growth"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawEPSEstChgforCurrentYear);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawEPSEstChgforCurrentYear() {

  "use strict";

  var raw = $("#eps-est-chg-for-current-year").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "EPS Est % Chg for Current Year",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("eps-est-chg-for-current-year"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawSalesChgLastQtr);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawSalesChgLastQtr() {

  "use strict";

  var raw = $("#sales-chg-last-qtr").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Sales % Chg (Last Qtr)",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("sales-chg-last-qtr"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawThreeYrSalesGrowthRate);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawThreeYrSalesGrowthRate() {

  "use strict";

  var raw = $("#three-yr-sales-growth-rate").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "3 Yr Sales Growth Rate",
    legend: {
      position: "none",
    }
  };


  var classicChart = new google.visualization.ColumnChart(document.getElementById("three-yr-sales-growth-rate"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawAnnualPreTaxMargin);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawAnnualPreTaxMargin() {

  "use strict";

  var raw = $("#annual-pretax-margin").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Annual Pre-Tax Margin",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("annual-pretax-margin"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawAnnualROE);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawAnnualROE() {

  "use strict";

  var raw = $("#annual-roe").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Annual ROE",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("annual-roe"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawDebtEquityRatio);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawDebtEquityRatio() {

  "use strict";

  var raw = $("#debt-equity-ratio").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Debt/Equity Ratio",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("debt-equity-ratio"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawOff52WeekHigh);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawOff52WeekHigh() {

  "use strict";

  var raw = $("#off-52-week-high").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "% Off 52 Week High",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("off-52-week-high"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawPricevs50DayMovingAverage);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawPricevs50DayMovingAverage() {

  "use strict";

  var raw = $("#price-vs-50-day-moving-average").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Price vs. 50-Day Moving Average",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("price-vs-50-day-moving-average"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawFiftyDayAverageVolume);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawFiftyDayAverageVolume() {

  "use strict";

  var raw = $("#fifty-day-average-volume").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "50-Day Average Volume",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("fifty-day-average-volume"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawChangeInFundsOwningStock);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawChangeInFundsOwningStock() {

  "use strict";

  var raw = $("#change-in-funds-owning-stock").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "% Change In Funds Owning Stock",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("change-in-funds-owning-stock"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

google.charts.setOnLoadCallback(drawQtrsOfIncreasingFundOwnership);

//////////////////////////////////////////////////////////////////////////////////////////////////////

function drawQtrsOfIncreasingFundOwnership() {

  "use strict";

  var raw = $("#qtrs-of-increasing-fund-ownership").data("ref");

  var body = JSON.parse(atob(raw));

  var data = google.visualization.arrayToDataTable(body);

  var classicOptions = {
    width: width,
    height: width / 2,
    hAxis: {
      slantedText: true,
    },
    title: "Qtrs Of Increasing Fund Ownership",
    legend: {
      position: "none",
    }
  };

  var classicChart = new google.visualization.ColumnChart(document.getElementById("qtrs-of-increasing-fund-ownership"));
  classicChart.draw(data, classicOptions);

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
