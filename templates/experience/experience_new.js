//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-note").focus();

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-create").click(function() {

  if (isProcessing()) {
    return;
  }

  var note = $("#input-note").val().trim();

  inProcess("#button-create");

  var exData = {
    "Data": note,
  };

  var exJsonBody = JSON.stringify(exData);

  $.ajax({
    type: "POST",
    url: "/experience",
    data: exJsonBody,
    success: function(data) {
      outProcess("#button-create");
      window.location = "/experience";
    },
    error: function(xhr, err) {
      outProcess("#button-create");
      alert(xhr.responseText);
    },
  });
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-cancel").click(function() {

  if (isProcessing()) {
    return;
  }

  window.location = "/experience";
});

//////////////////////////////////////////////////////////////////////////////////////////////////////
