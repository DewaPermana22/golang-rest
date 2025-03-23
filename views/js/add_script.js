document.getElementById("addProductForm").addEventListener("submit", async function (e) {
    e.preventDefault();

        setTimeout(() => {
        notification.classList.add("translate-y-10", "opacity-0");

        // Close modal
        setTimeout(() => {
            window.parent.postMessage("closeModal", "*");
        }, 500);
    }, 2000);

    const productName = document.getElementById('productName').value;
    const categoryId = document.getElementById("categoryList").value;
    const price = document.getElementById('productPrice').value;
    if (!productName || !categoryId || !price) {
        alert("Harap isi semua field!");
        return;
    }

    const requestData = {
        name: productName,
        category_id: parseInt(categoryId),
        price: parseFloat(price)
    };

    console.log("Data yang dikirim:", requestData);

    try {
        const response = await axios.post("http://localhost:5000/products/add", requestData, {
            headers: { "Content-Type": "application/json" }
        });
        console.log("Product added:", response.data);
        // Show
        const notification = document.getElementById("notification");
        notification.classList.remove("hidden");

        // Animasi
        setTimeout(() => {
            notification.classList.remove("translate-y-10", "opacity-0");
        }, 10);

        window.parent.postMessage("refreshProductList", "*");
    } catch (error) {
        console.error("Failed to add product:", error);
    }

});

window.addEventListener("message", async (event) => {
    if (event.data === "loadCategories") {
        await getCategory();
    }
});

const getCategory = async () => {
    try {
        const response = await axios.get("http://localhost:5000/categories");
        console.log("Data dari API:", response.data);
        const categories = response.data;
        const categoryList = document.getElementById("categoryList");

        categoryList.innerHTML = '<option value="">Pilih Kategori</option>';

        categories.forEach((category) => {
            const option = document.createElement("option");
            option.value = category.id;
            option.text = category.name;
            categoryList.appendChild(option);
        });
    } catch (error) {
        console.error(error);
    }
}