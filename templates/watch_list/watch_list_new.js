//////////////////////////////////////////////////////////////////////////////////////////////////////

//$("#input-password").focus();

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
  //if (isProcessing()) {
  //return;
  //}

  //window.location = "/action";
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