document.addEventListener("DOMContentLoaded", function () {
    getProducts();
});

window.addEventListener("message", (event) => {
    if (event.data === "refreshProductList") {
        console.log("Refreshing product list...");
        getProducts();
    }
});

let currentPage = 1
const limit = 10
const modal = document.getElementById("modal");
const modalFrame = document.getElementById("modalFrame");
const closeModal = document.getElementById("closeModal");
const addProductBtn = document.getElementById("addProductBtn");
const editButtons = document.querySelectorAll(".editBtn");
const tableContent = document.getElementById("productList");
const productCount = document.getElementById("productCount");

addProductBtn.addEventListener("click", () => {
    modalFrame.src = "/views/modals/add.html";
    modal.classList.remove("hidden");
    document.body.style.overflow = 'hidden';
});

// Open modal when "Edit" is clicked
editButtons.forEach((btn) => {
    btn.addEventListener("click", (e) => {
        const productId = e.target.getAttribute("data-id");
        modalFrame.src = `/views/modals/edit.html?id=${productId}`;
        modal.classList.remove("hidden");
        document.body.style.overflow = 'hidden';
    });
});

closeModal.addEventListener("click", () => {
    modal.classList.add("hidden");
    document.body.style.overflow = 'auto';
});

window.addEventListener("click", (e) => {
    if (e.target === modal) {
        modal.classList.add("hidden");
        document.body.style.overflow = 'auto';
    }
});

window.addEventListener('message', function (event) {
    if (event.data === 'closeModal') {
        modal.classList.add("hidden");
        document.body.style.overflow = 'auto';
    }
});

// Fetch product data


const getProducts = async (page = 1) => {
    await axios.get(`http://localhost:5000/products?page=${page}&limit=${limit}`)
        .then(response => {
            if (!response.data || !Array.isArray(response.data.data)) {
                console.error("Error: Data products tidak ditemukan!", response.data);
                return;
            }

            const products = response.data.data;
            const totalItems = products.length;
            tableContent.innerHTML = "";
            products.forEach((product, index) => {
                const row =
                    `
            <tr class="table-row hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">${index + 1}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">${product.name}</div>
                <div class="text-xs text-gray-500">SKU: PRD-A-001</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="badge bg-green-100 text-green-800">${product.category_name}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="font-medium">Rp ${product.price}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm">
                <button class="editBtn btn bg-amber-500 hover:bg-amber-600 text-white px-3 py-1.5 rounded-lg shadow-sm" data-id="1">
                  <i class="fas fa-edit mr-1"></i> Edit
                </button>
                <button class="btn bg-red-500 hover:bg-red-600 text-white px-3 py-1.5 rounded-lg shadow-sm ml-2">
                  <i class="fas fa-trash-alt mr-1"></i> Delete
                </button>
              </td>
            </tr>
            `;
                tableContent.innerHTML += row;
            });
            const count = products.length;
            productCount.textContent = `${count}`;
            updatePagination(page, limit, totalItems);
        })
        .catch(error => {
            console.log(error);
        });
}

modalFrame.addEventListener("load", () => {
    modalFrame.contentWindow.postMessage("loadCategories", "*");
});

const updatePagination = (page, limit, totalItems) => {
    const totalPages = Math.ceil(totalItems / limit);
    const paginationContainer = document.getElementById("paginationContainer");
    paginationContainer.innerHTML = ""; // Reset pagination sebelum update

    const paginationInfo = document.createElement("div");
    paginationInfo.classList.add("text-sm", "text-gray-500");
    paginationInfo.innerHTML = `
        Showing <span class="font-medium">${(page - 1) * limit + 1}-${Math.min(page * limit, totalItems)}</span> 
        of <span class="font-medium">${totalItems}</span> products
    `;
    paginationContainer.appendChild(paginationInfo);

    const paginationButtons = document.createElement("div");
    paginationButtons.classList.add("flex", "space-x-1");

    const prevButton = document.createElement("button");
    prevButton.classList.add("btn", "px-3", "py-1", "border", "border-gray-300", "rounded-lg", "text-gray-700", "bg-white", "hover:bg-gray-50");
    prevButton.innerHTML = `<i class="fas fa-chevron-left text-xs"></i>`;
    prevButton.disabled = page === 1;
    prevButton.addEventListener("click", () => changePage(page - 1));
    paginationButtons.appendChild(prevButton);

    for (let i = 1; i <= totalPages; i++) {
        const pageButton = document.createElement("button");
        pageButton.classList.add("btn", "px-3", "py-1", "border", ...(i === page ? ["bg-indigo-50", "border-indigo-300", "text-indigo-700", "font-medium"] : ["border-gray-300", "text-gray-700", "bg-white", "hover:bg-gray-50"]));
        pageButton.textContent = i;
        pageButton.addEventListener("click", () => changePage(i));
        paginationButtons.appendChild(pageButton);
    }

    const nextButton = document.createElement("button");
    nextButton.classList.add("btn", "px-3", "py-1", "border", "border-gray-300", "rounded-lg", "cursor-pointer", "text-gray-700", "bg-white", "hover:bg-gray-50");
    nextButton.innerHTML = `<i class="fas fa-chevron-right  text-xs"></i>`;
    nextButton.disabled = page === totalPages;
    nextButton.addEventListener("click", () => changePage(page + 1));
    paginationButtons.appendChild(nextButton);

    paginationContainer.appendChild(paginationButtons);
};


const changePage = (page) => {
    currentPage = page;
    getProducts(page);
};

