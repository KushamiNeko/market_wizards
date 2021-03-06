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

$("#button-flag").click(function() {
  flag = !flag;

  if (flag) {
    $("#button-flag-icon").html("flag");

    $("#button-flag").removeClass("btn-dark");
    $("#button-flag").addClass("btn-danger");

  } else {
    $("#button-flag-icon").html("outlined_flag");

    $("#button-flag").removeClass("btn-danger");
    $("#button-flag").addClass("btn-dark");
  }

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

  //if (!$("#dropdownMenu-priority").html().includes(":")) {
  //$("#validate-priority").show();
  //$("#dropdownMenu-priority").focus();
  //return;
  //}

  //var priority = $("#dropdownMenu-priority").html().split(":")[1].trim();

  if (!$("#dropdownMenu-status").html().includes(":")) {
    $("#validate-status").show();
    $("#dropdownMenu-status").focus();
    return;
  }

  var status = $("#dropdownMenu-status").html().split(":")[1].trim();

  if (!$("#dropdownMenu-grs").html().includes(":")) {
    $("#validate-grs").show();
    $("#dropdownMenu-grs").focus();
    return;
  }

  var grs = $("#dropdownMenu-grs").html().split(":")[1].trim();

  if (!$("#dropdownMenu-rs").html().includes(":")) {
    $("#validate-rs").show();
    $("#dropdownMenu-rs").focus();
    return;
  }

  var rs = $("#dropdownMenu-rs").html().split(":")[1].trim();

  if (!$("#dropdownMenu-fundamentals").html().includes(":")) {
    $("#validate-fundamentals").show();
    $("#dropdownMenu-fundamentals").focus();
    return;
  }

  var fundamentals = $("#dropdownMenu-fundamentals").html().split(":")[1].trim();

  //if ($("#input-note").val().match(/^[A-Z]+$/) === null) {
  //$("#validate-note").show();
  //$("#input-note").focus();
  //return;
  //}

  var note = $("#input-note").val().trim();

  inProcess("#button-create");

  var wlData = {
    "Symbol": symbol,
    "Price": price,
    //"Priority": priority,
    "GRS": grs,
    "RS": rs,
    "Fundamentals": fundamentals,
    "Status": status,
    "Note": note,
    "Flag": flag,
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
dropdownClick("#dropdownMenu-priority", "#dropdown-priority-e");
dropdownClick("#dropdownMenu-priority", "#dropdown-priority-p");

dropdownClick("#dropdownMenu-status", "#dropdown-status-a");
dropdownClick("#dropdownMenu-status", "#dropdown-status-b");
dropdownClick("#dropdownMenu-status", "#dropdown-status-c");
dropdownClick("#dropdownMenu-status", "#dropdown-status-d");
dropdownClick("#dropdownMenu-status", "#dropdown-status-e");

dropdownClick("#dropdownMenu-grs", "#dropdown-grs-a");
dropdownClick("#dropdownMenu-grs", "#dropdown-grs-b");
dropdownClick("#dropdownMenu-grs", "#dropdown-grs-c");

dropdownClick("#dropdownMenu-rs", "#dropdown-rs-a");
dropdownClick("#dropdownMenu-rs", "#dropdown-rs-b");
dropdownClick("#dropdownMenu-rs", "#dropdown-rs-c");
dropdownClick("#dropdownMenu-rs", "#dropdown-rs-d");

dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-a");
dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-b");
dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-c");
dropdownClick("#dropdownMenu-fundamentals", "#dropdown-fundamentals-d");

//////////////////////////////////////////////////////////////////////////////////////////////////////
