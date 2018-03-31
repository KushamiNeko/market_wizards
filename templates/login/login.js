//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#input-email").focus();

//////////////////////////////////////////////////////////////////////////////////////////////////////

function action(email, password, url, method, errcb) {

  $(".invalid-feedback").hide();

  if (isProcessing()) {
    return;
  }

  if (!validateEmail(email)) {
    $("#validate-email").show();
    return;
  }

  if (password === "") {
    $("#validate-password").show();
    return;
  }

  inProcess("#button-login");

  var data = {};
  data.Email = email;
  data.Password = password;

  var jsonBody = JSON.stringify(data);

  $.ajax({
    type: method,
    url: url,
    data: jsonBody,
    success: function(data) {
      window.location = data;

      outProcess();
    },
    error: function(xhr, err) {
      errcb();

      outProcess("#button-login");
    },
  });
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-create").click(function() {
  action(
    $("#input-email").val() + "@gmail.com",
    $("#input-password").val(),
    "/user",
    "POST",
    function() {
      $("#validate-email").show();
    });
});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-login").click(function() {
  action(
    $("#input-email").val() + "@gmail.com",
    $("#input-password").val(),
    "/login",
    "POST",
    function() {
      $("#validate-email").show();
      $("#validate-password").show();
    });

});

//////////////////////////////////////////////////////////////////////////////////////////////////////

$("#button-reset").click(function() {
  action(
    $("#input-email").val() + "@gmail.com",
    "0000",
    "/user",
    "PUT",
    function() {
      $("#validate-password").show();
    });
});

//////////////////////////////////////////////////////////////////////////////////////////////////////
