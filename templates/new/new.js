//////////////////////////////////////////////////////////////////////////////////////////////////////

$(".invalid-feedback").hide();

//////////////////////////////////////////////////////////////////////////////////////////////////////

buyOrderInputs();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-back").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/action";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#new-modal-body").on("input", function() {
  this.style.height = "auto";
  this.style.height = (this.scrollHeight) + "px";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#new-modal").on("shown.bs.modal", function() {
  $("#new-modal-body").trigger("input");
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
    buyOrderInputs();
  }

  if (sell) {
    $("#button-type").html("SELL");
    sellOrderInputs();
  }

  $("#input-revenue").val("");
  $("#input-cost").val("");
  $("#input-gainD").val("");
  $("#input-gainP").val("");
  $("#input-daysheld").val("");
  $("#input-date-of-purchase").val("");

  $("#dropdownMenu").html("Buy Point");
  buyPoint = "";

  calculate();
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

function buyOrderInputs() {
  $("#input-revenue").attr("disabled", "disabled");
  $("#input-daysheld").attr("disabled", "disabled");
  $("#input-date-of-purchase").attr("disabled", "disabled");

  $("#input-cost").removeAttr("disabled");
  $("#input-stage").removeAttr("disabled");
  $("#dropdownMenu").removeAttr("disabled");
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

function sellOrderInputs() {
  $("#input-cost").attr("disabled", "disabled");
  $("#input-stage").attr("disabled", "disabled");
  $("#dropdownMenu").attr("disabled", "disabled");

  $("#input-revenue").removeAttr("disabled");
  $("#input-daysheld").removeAttr("disabled");
  $("#input-date-of-purchase").removeAttr("disabled");
}

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

function inputPasteCleanup(jqueryID) {
  $(jqueryID).val($(jqueryID).val().replace(",", "").trim());
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-price").change(function() {
  inputPasteCleanup("#input-price");
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-revenue").change(function() {
  inputPasteCleanup("#input-revenue");
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-cost").change(function() {
  inputPasteCleanup("#input-cost");
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-share").change(function() {
  if (sell) {
    $("#input-date-of-purchase").trigger("change");
  }
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-date-of-purchase").change(function() {
  $(".invalid-feedback").hide();

  if ($("#input-date-of-purchase").val().match(/^\d{8}$/) === null) {
    $("#validate-date-of-purchase").show();
    $("#input-date-of-purchase").focus();
    return;
  }

  var dateOfPurchase = $("#input-date-of-purchase").val();

  if ($("#input-symbol").val().toUpperCase().match(/^[A-Z]+$/) === null) {
    $("#validate-symbol").show();
    $("#input-symbol").focus();
    return;
  }

  var symbol = $("#input-symbol").val().toUpperCase();

  if ($("#input-share").val().match(/^\d+$/) === null) {
    $("#validate-share").show();
    $("#input-share").focus();
    return;
  }

  var shareSold = parseInt($("#input-share").val());

  $.ajax({
    type: "Get",
    url: "/transaction?DateOfPurchase=" + encodeURIComponent(dateOfPurchase) + "&Symbol=" + encodeURIComponent(symbol),
    success: function(data) {

      var cost = parseFloat(data.Cost);

      var sharePurchase = parseFloat(data.Share);

      var newCost = cost * (shareSold / sharePurchase);

      newCost = newCost.toFixed(decimal);

      $("#input-cost").val(newCost);

      $("#input-stage").val(data.Stage);
      $("#dropdownMenu").html(data.BuyPoint);

      calculate();
    },
    error: function(xhr, err) {
      //$("#new-modal-body").val(xhr.responseText);
      //$("#new-modal").modal("show");

      alert(xhr.responseText);
    },
  });
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

var transactionStatus = 0;

//////////////////////////////////////////////////////////////////////////////////////////////////////

function transactionOK(newLocation) {

  transactionStatus += 1;

  if (transactionStatus === 3) {

    transactionStatus = 0;

    outProcess("#button-create");

    window.location = newLocation;

  }
}

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

  if (buy) {
    if (buyPoint === "") {
      $("#validate-buypoint").show();
      $("#dropdownMenu").focus();
      return;
    }
  }

  if (sell) {
    if ($("#input-revenue").val().match(/^[0-9.]+$/) === null) {
      $("#validate-revenue").show();
      $("#input-revenue").focus();
      return;
    }
  }

  var revenue = parseFloat($("#input-revenue").val());

  if (sell) {
    if ($("#input-date-of-purchase").val().match(/^\d{8}$/) === null) {
      $("#validate-date-of-purchase").show();
      $("#input-date-of-purchase").focus();
      return;
    }
  }

  var dateOfPurchase = parseInt($("#input-date-of-purchase").val());

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
    if ($("#input-daysheld").val().match(/^[0-9]+$/) === null) {
      $("#validate-daysheld").show();
      $("#input-daysheld").focus();
      return;
    }
  }

  var daysheld = parseInt($("#input-daysheld").val());

  var note = $("#input-note").val();

  if ($("#img-ibd").attr("src") === undefined) {
    $("#validate-img-ibd").show();
    $("#button-img-ibd").focus();
    return;
  }

  if ($("#img-marketsmith").attr("src") === undefined) {
    $("#validate-img-marketsmith").show();
    $("#button-img-marketsmith").focus();
    return;
  }

  //if ($("#img-d").attr("src") === undefined) {
  //$("#validate-img-d").show();
  //$("#button-img-d").focus();
  //return;
  //}

  //if ($("#img-w").attr("src") === undefined) {
  //$("#validate-img-w").show();
  //$("#button-img-w").focus();
  //return;
  //}

  //if ($("#img-ndqc-d").attr("src") === undefined) {
  //$("#validate-img-ndqc-d").show();
  //$("#button-img-ndqc-d").focus();
  //return;
  //}

  //if ($("#img-ndqc-w").attr("src") === undefined) {
  //$("#validate-img-ndqc-w").show();
  //$("#button-img-ndqc-w").focus();
  //return;
  //}

  //if ($("#img-sp5-d").attr("src") === undefined) {
  //$("#validate-img-sp5-d").show();
  //$("#button-img-sp5-d").focus();
  //return;
  //}

  //if ($("#img-sp5-w").attr("src") === undefined) {
  //$("#validate-img-sp5-w").show();
  //$("#button-img-sp5-w").focus();
  //return;
  //}

  //if ($("#img-nyc-d").attr("src") === undefined) {
  //$("#validate-img-nyc-d").show();
  //$("#button-img-nyc-d").focus();
  //return;
  //}

  //if ($("#img-nyc-w").attr("src") === undefined) {
  //$("#validate-img-nyc-w").show();
  //$("#button-img-nyc-w").focus();
  //return;
  //}

  //if ($("#img-djia-d").attr("src") === undefined) {
  //$("#validate-img-djia-d").show();
  //$("#button-img-djia-d").focus();
  //return;
  //}

  //if ($("#img-djia-w").attr("src") === undefined) {
  //$("#validate-img-djia-w").show();
  //$("#button-img-djia-w").focus();
  //return;
  //}

  //if ($("#img-rus-d").attr("src") === undefined) {
  //$("#validate-img-rus-d").show();
  //$("#button-img-rus-d").focus();
  //return;
  //}

  //if ($("#img-rus-w").attr("src") === undefined) {
  //$("#validate-img-rus-w").show();
  //$("#button-img-rus-w").focus();
  //return;
  //}

  inProcess("#button-create");

  var ibdData = {
    "Date": date,
    "Symbol": symbol,
    //"Object": $("#img-ibd").attr("src"),
    "Data": $("#img-ibd").attr("src"),
  };

  var ibdJsonBody = JSON.stringify(ibdData);

  $.ajax({
    type: "POST",
    url: "/ibd",
    data: ibdJsonBody,
    success: function(data) {
      //outProcess("#button-create");
      transactionOK("/action");
    },
    error: function(xhr, err) {
      //$("#new-modal-body").val(xhr.responseText);
      //$("#new-modal").modal("show");

      //outProcess("#button-create");

      alert(xhr.responseText);
    },
  });

  var marketsmithData = {
    "Date": date,
    "Symbol": symbol,
    //"Object": $("#img-marketsmith").attr("src"),
    "Data": $("#img-marketsmith").attr("src"),
  };

  var marketsmithJsonBody = JSON.stringify(marketsmithData);

  $.ajax({
    type: "POST",
    url: "/marketsmith",
    data: marketsmithJsonBody,
    success: function(data) {
      //outProcess("#button-create");
      transactionOK("/action");
    },
    error: function(xhr, err) {
      //$("#new-modal-body").val(xhr.responseText);
      //$("#new-modal").modal("show");

      //outProcess("#button-create");
      alert(xhr.responseText);
    },
  });


  var data = {
    //"Order": "buy",
    "Date": date,
    "Symbol": symbol,
    "Price": price,
    "Share": share,
    //"BuyPoint": buyPoint,
    //"Cost": cost,
    //"Stage": stage,
    "Note": note,
    //"JsonIBDCheckup": $("#img-ibd").attr("src"),
    //"JsonChartD": $("#img-d").attr("src"),
    //"JsonChartW": $("#img-w").attr("src"),
    //"JsonChartNDQCD": $("#img-ndqc-d").attr("src"),
    //"JsonChartNDQCW": $("#img-ndqc-w").attr("src"),
    //"JsonChartSP5D": $("#img-sp5-d").attr("src"),
    //"JsonChartSP5W": $("#img-sp5-w").attr("src"),
    //"JsonChartNYCD": $("#img-nyc-d").attr("src"),
    //"JsonChartNYCW": $("#img-nyc-w").attr("src"),
    //"JsonChartDJIAD": $("#img-djia-d").attr("src"),
    //"JsonChartDJIAW": $("#img-djia-w").attr("src"),
    //"JsonChartRUSD": $("#img-rus-d").attr("src"),
    //"JsonChartRUSW": $("#img-rus-w").attr("src"),
  };

  if (buy) {
    data.Order = "buy";
    data.BuyPoint = buyPoint;
    data.Cost = cost;
    data.Stage = stage;
  }

  if (sell) {
    data.Order = "sell";
    data.DateOfPurchase = dateOfPurchase;
    data.Revenue = revenue;
    data.Cost = cost;
    data.GainD = gainD;
    data.GainP = gainP;
    data.DaysHeld = daysheld;
  }

  var jsonBody = JSON.stringify(data);

  var order = "buy";

  if (sell) {
    order = "sell";
  }

  $.ajax({
    type: "POST",
    url: "/transaction?Order=" + encodeURIComponent(order),
    data: jsonBody,
    success: function(data) {
      //outProcess("#button-create");
      //window.location = data;
      transactionOK("/action");
    },
    error: function(xhr, err) {
      //$("#new-modal-body").val(xhr.responseText);
      //$("#new-modal").modal("show");

      //outProcess("#button-create");
      alert(xhr.responseText);
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

//imgChange("d");
//imgChange("w");

//imgChange("ndqc-d");
//imgChange("ndqc-w");

//imgChange("sp5-d");
//imgChange("sp5-w");

//imgChange("nyc-d");
//imgChange("nyc-w");

//imgChange("djia-d");
//imgChange("djia-w");

//imgChange("rus-d");
//imgChange("rus-w");

imgChange("ibd");
imgChange("marketsmith");

//////////////////////////////////////////////////////////////////////////////////////////////////////

function dropdownClick(target) {
  $(target).click(function() {
    $("#dropdownMenu").html($(target).html());
    buyPoint = $(target).html().trim();
  });
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

dropdownClick("#dropdown-vcp-early-entry");
dropdownClick("#dropdown-earnings-report");
dropdownClick("#dropdown-gap");
dropdownClick("#dropdown-tennis-ball");
dropdownClick("#dropdown-pullback");
dropdownClick("#dropdown-trend-line");
dropdownClick("#dropdown-resistance");
dropdownClick("#dropdown-support");
dropdownClick("#dropdown-cup-with-handle");
dropdownClick("#dropdown-cup");
dropdownClick("#dropdown-double-bottom");
dropdownClick("#dropdown-flat-base");
dropdownClick("#dropdown-ipo-base");
dropdownClick("#dropdown-tight-area");
dropdownClick("#dropdown-consolidation");
dropdownClick("#dropdown-21d-pullback");
dropdownClick("#dropdown-50d-pullback");
dropdownClick("#dropdown-new-high");
dropdownClick("#dropdown-old-entry");

//////////////////////////////////////////////////////////////////////////////////////////////////////
