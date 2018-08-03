//////////////////////////////////////////////////////////////////////////////////////////////////////

$(".invalid-feedback").hide();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-new").focus();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-back").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/action";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-new").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/watchlist?Action=new";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-calculate").click(function() {
  $(".invalid-feedback").hide();

  if (isProcessing()) {
    return;
  }

  if ($("#input-capital").val() !== "") {
    if ($("#input-capital").val().match(/^[0-9.]+$/) === null) {
      $("#validate-capital").show();
      $("#input-capital").focus();
      return;
    }
  }

  if ($("#input-size").val() !== "") {
    if ($("#input-size").val().match(/^[0-9.]+$/) === null) {
      $("#validate-size").show();
      $("#input-size").focus();
      return;
    }
  }

  if ($("#input-symbol").val() !== "") {
    if ($("#input-symbol").val().toUpperCase().match(/^[A-Z]+$/) === null) {
      $("#validate-symbol").show();
      $("#input-symbol").focus();
      return;
    }

    var symbol = $("#input-symbol").val().toUpperCase();

    window.location = "/watchlist?Capital=" + $("#input-capital").val() + "&Size=" + $("#input-size").val() + "&Symbol=" + symbol;
  } else {

    window.location = "/watchlist?Capital=" + $("#input-capital").val() + "&Size=" + $("#input-size").val();
  }

});

//////////////////////////////////////////////////////////////////////////////////////////////////////
