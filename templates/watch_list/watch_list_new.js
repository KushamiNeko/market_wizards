//////////////////////////////////////////////////////////////////////////////////////////////////////

$(".invalid-feedback").hide();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-symbol").focus();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-back").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/action";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-cancel").click(function() {
  if (isProcessing()) {
    return;
  }

  window.location = "/watchlist";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-create").click(function() {
  $(".invalid-feedback").hide();

  if (isProcessing()) {
    return;
  }

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

  if (!$("#dropdownMenu-priority").html().includes(":")) {
    $("#validate-priority").show();
    $("#dropdownMenu-priority").focus();
    return;
  }

  var priority = $("#dropdownMenu-priority").html().split(":")[1].trim();

  if (!$("#dropdownMenu-status").html().includes(":")) {
    $("#validate-status").show();
    $("#dropdownMenu-status").focus();
    return;
  }

  var status = $("#dropdownMenu-status").html().split(":")[1].trim();

  if (!$("#dropdownMenu-fundamentals").html().includes(":")) {
    $("#validate-fundamentals").show();
    $("#dropdownMenu-fundamentals").focus();
    return;
  }

  var fundamentals = $("#dropdownMenu-fundamentals").html().split(":")[1].trim();

  inProcess("#button-create");

  var wlData = {
    "Symbol": symbol,
    "Price": price,
    "Priority": priority,
    "Fundamentals": fundamentals,
    "Status": status,
  };

  var wlJsonBody = JSON.stringify(wlData);

  $.ajax({
    type: "POST",
    url: "/watchlist",
    data: wlJsonBody,
    success: function(data) {
      outProcess("#button-create");
      window.location = "/watchlist";
    },
    error: function(xhr, err) {
      outProcess("#button-create");
      alert(xhr.responseText);
    },
  });

});

////////////////////////////////////////////////////////////////////////////////////////////////////////

function dropdownClick(menu, target) {
  $(target).click(function() {

    var label = $(menu).html();

    if (label.includes(":")) {
      label = label.split(":")[0].trim();
    }

    $(menu).html(label + " : " + $(target).html());
  });
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

dropdownClick("#dropdownMenu-priority", "#dropdown-priority-a");
dropdownClick("#dropdownMenu-priority", "#dropdown-priority-b");
dropdownClick("#dropdownMenu-priority", "#dropdown-priority-c");

dropdownClick("#dropdownMenu-status", "#dropdown-status-a");
dropdownClick("#dropdownMenu-status", "#dropdown-status-b");

dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-a");
dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-b");
dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-c");

//////////////////////////////////////////////////////////////////////////////////////////////////////
