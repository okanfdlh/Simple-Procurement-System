const API_URL = "http://localhost:3000/api";

function apiRequest(method, url, data, onSuccess, onError) {
  $("#loading").removeClass("hidden").addClass("flex");

  $.ajax({
    method,
    url: API_URL + url,
    contentType: "application/json",
    data: data ? JSON.stringify(data) : null,
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token")
    },
    success: function (res) {
      $("#loading").addClass("hidden").removeClass("flex");
      onSuccess(res);
    },
    error: function (xhr) {
      $("#loading").addClass("hidden").removeClass("flex");
      onError && onError(xhr.responseJSON);
    }
  });
}

