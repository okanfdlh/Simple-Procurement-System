let cart = [];

/* =======================
   LOAD DATA
======================= */
apiRequest("GET", "/suppliers", null, res => {
  res.data.forEach(s => {
    $("#supplier").append(
      `<option value="${s.id}">${s.name}</option>`
    );
  });
});

apiRequest("GET", "/items", null, res => {
  res.data.forEach(i => {
    $("#item").append(
      `<option value="${i.id}">${i.name}</option>`
    );
  });
});

/* =======================
   CART LOGIC
======================= */
$("#addItem").click(function () {
  const itemId = parseInt($("#item").val());
  const qty = parseInt($("#qty").val());

  if (!itemId || qty <= 0) {
    alert("Item & qty wajib diisi");
    return;
  }

  cart.push({ item_id: itemId, qty });
  renderCart();
});

function renderCart() {
  $("#cart").html("");

  cart.forEach((item, index) => {
    $("#cart").append(`
      <li class="flex justify-between items-center border-b py-1">
        <span>Item ${item.item_id} - Qty ${item.qty}</span>
        <button
          class="text-red-500 text-sm"
          onclick="removeItem(${index})">
          Remove
        </button>
      </li>
    `);
  });
}

function removeItem(index) {
  cart.splice(index, 1);
  renderCart();
}

/* =======================
   SUBMIT PURCHASE
======================= */
$("#submit").click(function () {
  const supplierId = parseInt($("#supplier").val());

  if (!supplierId) {
    alert("Supplier wajib dipilih");
    return;
  }

  if (cart.length === 0) {
    alert("Cart masih kosong");
    return;
  }

  apiRequest(
    "POST",
    "/purchasing",
    {
      supplier_id: supplierId,
      items: cart
    },
    function () {
      alert("Purchase created");
      cart = [];
      window.location.href = "dashboard.html";
    },
    function (err) {
      alert(err.message || "Failed to create purchase");
      console.error(err);
    }
  );
});
