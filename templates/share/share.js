function validateEmail(email) {
  var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(email);
}

function fileChanged(input, img) {
  if (!FileReader) {
    console.log('FileReader unsupport');
    return;
  }

  if (input[0].files && input[0].files[0]) {

    var reader = new FileReader();

    reader.onload = function(e) {

      if (img !== null) {
        img.attr("src", e.target.result);
      }

      //this.content[i] = e.target.result.split(",", 2)[1];
    };

    reader.readAsDataURL(input[0].files[0]);
  }

}

function extractImg(img) {
  return img.split(",", 2)[1];
}

var processing = false;

function inProcess(target) {
  //console.log(target);
  processing = true;

  $(target).removeClass("btn-dark");
  $(target).addClass("btn-danger");
}

function outProcess(target) {
  processing = false;

  $(target).removeClass("btn-danger");
  $(target).addClass("btn-dark");
}

function isProcessing() {
  return processing;
}
