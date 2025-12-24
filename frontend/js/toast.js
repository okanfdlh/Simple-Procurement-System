function showToast(message, type = "success") {
  const toast = $("#toast");
  toast
    .removeClass()
    .addClass(`fixed top-5 right-5 px-4 py-2 rounded shadow text-white 
      ${type === "success" ? "bg-green-600" : "bg-red-600"}`)
    .text(message)
    .fadeIn();

  setTimeout(() => toast.fadeOut(), 3000);
}
