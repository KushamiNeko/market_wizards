//////////////////////////////////////////////////////////////////////////////////////////////////////

$(".invalid-feedback").hide();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-back").click(function() {
  window.location = "/action";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

var decimal = 2;

var buy = true;
var sell = false;

var buyPoint = "";

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-type").click(function() {
  buy = !buy;
  sell = !buy;

  if (buy) {
    $("#button-type").html("BUY");
    $("#input-revenue").attr("disabled", "disabled");
    $("#input-dayhold").attr("disabled", "disabled");
  }

  if (sell) {
    $("#button-type").html("SELL");
    $("#input-revenue").removeAttr("disabled");
    $("#input-dayhold").removeAttr("disabled");
  }

  $("#input-revenue").val("");
  $("#input-cost").val("");
  $("#input-gainD").val("");
  $("#input-gainP").val("");
  $("#input-dayhold").val("");

  calculate();
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

function calculate() {
  var revenue = 0;
  var cost = 0;
  var gainD = 0;
  var gainP = 0;

  if (sell) {
    if ($("#input-revenue").val().match(/^[0-9.]+$/) === null) {
      return;
    }

    if ($("#input-cost").val().match(/^[0-9.]+$/) === null) {
      return;
    }

    revenue = parseFloat($("#input-revenue").val());
    cost = parseFloat($("#input-cost").val());

    gainD = (revenue - cost).toFixed(decimal);
    $("#input-gainD").val(gainD.toString());

    gainP = ((gainD / cost) * 100).toFixed(decimal);

    $("#input-gainP").val(gainP.toString());
  }

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-revenue").focusout(function() {
  calculate();
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-cost").focusout(function() {
  calculate();
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-create").click(function() {
  $(".invalid-feedback").hide();

  if (isProcessing()) {
    return;
  }

  if ($("#input-date").val().match(/^\d{8}$/) === null) {
    $("#validate-date").show();
    $("#input-date").focus();
    return;
  }

  var date = parseInt($("#input-date").val());

  if ($("#input-symbol").val().toUpperCase().match(/^[A-Z]+$/) === null) {
    $("#validate-symbol").show();
    $("#input-symbol").focus();
    return;
  }

  var symbol = $("#input-symbol").val().toUpperCase();

  if ($("#input-price").val().match(/^[0-9.]+$/) === null) {
    $("#validate-price").show();
    $("#input-price").focus();
    return;
  }

  var price = parseFloat($("#input-price").val());

  if ($("#input-share").val().match(/^\d+$/) === null) {
    $("#validate-share").show();
    $("#input-share").focus();
    return;
  }

  var share = parseInt($("#input-share").val());

  if (sell) {
    if ($("#input-revenue").val().match(/^[0-9.]+$/) === null) {
      $("#validate-revenue").show();
      $("#input-revenue").focus();
      return;
    }
  }

  if (buyPoint === "") {
    $("#validate-buypoint").show();
    $("#dropdownMenu").focus();
    return;
  }

  var revenue = parseFloat($("#input-revenue").val());

  if ($("#input-cost").val().match(/^[0-9.]+$/) === null) {
    $("#validate-cost").show();
    $("#input-cost").focus();
    return;
  }

  var cost = parseFloat($("#input-cost").val());

  var gainD = parseFloat($("#input-gainD").val());

  var gainP = parseFloat($("#input-gainP").val());

  if ($("#input-stage").val().match(/^[0-9.]+$/) === null) {
    $("#validate-stage").show();
    $("#input-stage").focus();
    return;
  }

  var stage = parseFloat($("#input-stage").val());

  if (sell) {
    if ($("#input-dayhold").val().match(/^[0-9]+$/) === null) {
      $("#validate-dayhold").show();
      $("#input-dayhold").focus();
      return;
    }
  }

  var dayhold = parseInt($("#input-dayhold").val());

  var note = $("#input-note").val();

  if ($("#img-ibd").attr("src") === undefined) {
    $("#validate-img-ibd").show();
    $("#button-img-ibd").focus();
    return;
  }

  if ($("#img-d").attr("src") === undefined) {
    $("#validate-img-d").show();
    $("#button-img-d").focus();
    return;
  }

  if ($("#img-w").attr("src") === undefined) {
    $("#validate-img-w").show();
    $("#button-img-w").focus();
    return;
  }

  if ($("#img-ndqc-d").attr("src") === undefined) {
    $("#validate-img-ndqc-d").show();
    $("#button-img-ndqc-d").focus();
    return;
  }

  if ($("#img-ndqc-w").attr("src") === undefined) {
    $("#validate-img-ndqc-w").show();
    $("#button-img-ndqc-w").focus();
    return;
  }

  if ($("#img-sp5-d").attr("src") === undefined) {
    $("#validate-img-sp5-d").show();
    $("#button-img-sp5-d").focus();
    return;
  }

  if ($("#img-sp5-w").attr("src") === undefined) {
    $("#validate-img-sp5-w").show();
    $("#button-img-sp5-w").focus();
    return;
  }

  if ($("#img-nyc-d").attr("src") === undefined) {
    $("#validate-img-nyc-d").show();
    $("#button-img-nyc-d").focus();
    return;
  }

  if ($("#img-nyc-w").attr("src") === undefined) {
    $("#validate-img-nyc-w").show();
    $("#button-img-nyc-w").focus();
    return;
  }

  if ($("#img-djia-d").attr("src") === undefined) {
    $("#validate-img-djia-d").show();
    $("#button-img-djia-d").focus();
    return;
  }

  if ($("#img-djia-w").attr("src") === undefined) {
    $("#validate-img-djia-w").show();
    $("#button-img-djia-w").focus();
    return;
  }

  if ($("#img-rus-d").attr("src") === undefined) {
    $("#validate-img-rus-d").show();
    $("#button-img-rus-d").focus();
    return;
  }

  if ($("#img-rus-w").attr("src") === undefined) {
    $("#validate-img-rus-w").show();
    $("#button-img-rus-w").focus();
    return;
  }

  inProcess("#button-create");

  var data = {
    "Date": date,
    "Symbol": symbol,
    "Order": "buy",
    "Price": price,
    "Share": share,
    "BuyPoint": buyPoint,
    "Cost": cost,
    "Stage": stage,
    "Note": note,
    "JsonChartD": $("#img-d").attr("src"),
    "JsonChartW": $("#img-w").attr("src"),
    "JsonChartNDQCD": $("#img-ndqc-d").attr("src"),
    "JsonChartNDQCW": $("#img-ndqc-w").attr("src"),
    "JsonChartSP5D": $("#img-sp5-d").attr("src"),
    "JsonChartSP5W": $("#img-sp5-w").attr("src"),
    "JsonChartNYCD": $("#img-nyc-d").attr("src"),
    "JsonChartNYCW": $("#img-nyc-w").attr("src"),
    "JsonChartDJIAD": $("#img-djia-d").attr("src"),
    "JsonChartDJIAW": $("#img-djia-w").attr("src"),
    "JsonChartRUSD": $("#img-rus-d").attr("src"),
    "JsonChartRUSW": $("#img-rus-w").attr("src"),
    "JsonIBDCheckup": $("#img-ibd").attr("src"),
  };

  if (sell) {
    data.Order = "sell";
    data.Revenue = revenue;
    data.GainD = gainD;
    data.GainP = gainP;
    data.DayHold = dayhold;
  }

  var jsonBody = JSON.stringify(data);

  processing = true;

  $.ajax({
    type: "POST",
    url: "/transaction",
    data: jsonBody,
    success: function(data) {
      window.location = data;

      outProcess();
    },
    error: function(xhr, err) {
      console.log(err);

      outProcess("#button-create");
    },
  });

});

//////////////////////////////////////////////////////////////////////////////////////////////////////

function imgChange(target) {
  var button = "#button-img-" + target;
  var inputs = "#input-img-" + target;
  var img = "#img-" + target;

  $(button).click(function() {
    $(inputs).click();
  });

  $(inputs).change(function() {
    fileChanged($(inputs), $(img));
  });
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

imgChange("d");
imgChange("w");

imgChange("ndqc-d");
imgChange("ndqc-w");

imgChange("sp5-d");
imgChange("sp5-w");

imgChange("nyc-d");
imgChange("nyc-w");

imgChange("djia-d");
imgChange("djia-w");

imgChange("rus-d");
imgChange("rus-w");

imgChange("ibd");

//////////////////////////////////////////////////////////////////////////////////////////////////////

//function dropdownClick(target, value) {
function dropdownClick(target) {
  $(target).click(function() {
    $("#dropdownMenu").html($(target).html());
    //buyPoint = value;
    buyPoint = $(target).html();
  });
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

//dropdownClick("#dropdown-cup-with-handle", "Cup With Handle");
//dropdownClick("#dropdown-cup", "Cup");
//dropdownClick("#dropdown-double-bottom", "Double Bottom");
//dropdownClick("#dropdown-flat-base", "Flat Base");
//dropdownClick("#dropdown-ipo-base", "IPO Base");
//dropdownClick("#dropdown-tight-area", "Tight Area");
//dropdownClick("#dropdown-channel", "Channel");
//dropdownClick("#dropdown-consolidation", "Consolidation");
//dropdownClick("#dropdown-21d-pullback", "21D PullBack");
//dropdownClick("#dropdown-50d-pullback", "50D PullBack");
//dropdownClick("#dropdown-new-high", "New High");

dropdownClick("#dropdown-vcp-early-entry");
dropdownClick("#dropdown-earnings-report");
dropdownClick("#dropdown-gap");
dropdownClick("#dropdown-po-so");
dropdownClick("#dropdown-trend-line");
dropdownClick("#dropdown-resistance");
dropdownClick("#dropdown-support");
dropdownClick("#dropdown-cup-with-handle");
dropdownClick("#dropdown-cup");
dropdownClick("#dropdown-double-bottom");
dropdownClick("#dropdown-flat-base");
dropdownClick("#dropdown-ipo-base");
dropdownClick("#dropdown-tight-area");
//dropdownClick("#dropdown-channel");
dropdownClick("#dropdown-consolidation");
dropdownClick("#dropdown-21d-pullback");
dropdownClick("#dropdown-50d-pullback");
dropdownClick("#dropdown-new-high");

//"VCP Early Entry",
//"VCP Pivot",
//"Earnings Report",
//"Gap",
//"PO / SO",
//"Trend Line",
//"IPO Base",
//"Cup With Handle",
//"Cup",
//"Flat Base",
//"Tight Area",
//"Double Bottom",
//"Consolidation",
//"21D Pullback",
//"50D Pullback",

//////////////////////////////////////////////////////////////////////////////////////////////////////
