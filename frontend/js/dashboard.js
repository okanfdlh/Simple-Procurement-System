apiRequest("GET", "/items", null, function (res) {
  res.data.forEach(item => {
    $("#items").append(`
      <tr class="border-t">
        <td class="p-2">${item.name}</td>
        <td class="p-2 text-center">${item.stock}</td>
        <td class="p-2 text-center">${item.price}</td>
      </tr>
    `);
  });
});
