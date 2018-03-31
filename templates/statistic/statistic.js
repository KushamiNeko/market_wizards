//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-start-date").focus();

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

    window.location = "/statistic?start=" + $("#input-start-date").val() + "&end=" + $("#input-end-date").val();

  }

  outProcess("#button-query");

});

//////////////////////////////////////////////////////////////////////////////////////////////////////
