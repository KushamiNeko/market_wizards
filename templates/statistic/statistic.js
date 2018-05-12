//////////////////////////////////////////////////////////////////////////////////////////////////////

var infoIndex = 0;
var infoType = ["Info General", "Chart General", "Chart IBD", "Chart MarketSmith"];
var infoContainer = ["#info-general", "#chart-general", "#chart-ibd", "#chart-marketsmith"];

buttonTypeUpdate();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-type").click(function() {
  infoIndex += 1;
  infoIndex = infoIndex % infoType.length;

  buttonTypeUpdate();
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

function buttonTypeUpdate() {
  $("#button-type").html(infoType[infoIndex]);

  for (var i = 0; i < infoType.length; i++) {
    if (i === infoIndex) {
      $(infoContainer[i]).show();
    } else {
      $(infoContainer[i]).hide();
    }
  }
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-start-date").focus();

//$("#input-threshold").val("1.0");

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-back").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/action";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-query").click(function() {

  $(".invalid-feedback").hide();

  if (isProcessing()) {
    return;
  }

  inProcess("#button-query");

  if ($("#input-start-date").val() === "" && $("#input-end-date").val() === "") {

    window.location = "/statistic";

  } else {

    if ($("#input-start-date").val().match(/^\d{8}$/) === null) {
      $("#validate-start-date").show();
      $("#input-start-date").focus();

      outProcess("#button-query");
      return;
    }

    if ($("#input-end-date").val().match(/^\d{8}$/) === null) {
      $("#validate-end-date").show();
      $("#input-end-date").focus();

      outProcess("#button-query");
      return;
    }

    if ($("#input-threshold").val().match(/^[0-9.]+$/) === null) {
      $("#validate-threshold").show();
      $("#input-threshold").focus();

      outProcess("#button-query");
      return;
    }

    window.location =
      "/statistic?start=" + $("#input-start-date").val() + "&end=" + $("#input-end-date").val() +
      "&threshold=" + $("#input-threshold").val();
  }

  outProcess("#button-query");

});

//////////////////////////////////////////////////////////////////////////////////////////////////////
