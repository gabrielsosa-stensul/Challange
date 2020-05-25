const API_URL = "/api"

// listItems display all items in a sortable list.
const listItems = () => {
    $.ajax({
        url: API_URL + "/items/",
        type: 'GET',
        contentType: false,
        processData: false,
        success: function(data) {
            list.data.items = data && data.length ? data : []
            list.render();
            createSortableList()
        },
        error: function(data) {
            console.error(data)
        }
    });
};

// newItem display a popup to create a new item.
const newItem = () => {
    form.data.item = null;
    form.render();
    $('#form-component').modal('show');
}

// createItem post the data of a new item in the API.
const createItem = (e) => {
    e.preventDefault();
    formData = new FormData(e.target);
    $.ajax({
        url: API_URL + "/items/",
        type: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        success: function(data) {
            listItems()
            $('#form-component').modal('hide');
        },
        error: function(data) {
            showErrorMessages(data)
        }
    });
}

// editItem display a popup to update an existing item.
const editItem = id => {
    $.ajax({
        url: API_URL + "/items/" + id,
        type: 'GET',
        contentType: false,
        processData: false,
        success: function(data) {
            form.data.item = data;
            form.render();
            $('#form-component').modal('show');
        },
        error: function(data) {
            console.error(data)
        }
    });
};

// updateItem patch the data of an existing item in the API.
const updateItem = (e, id) => {
    e.preventDefault();
    formData = new FormData(e.target);
    $.ajax({
        url: API_URL + "/items/" + id,
        type: 'PATCH',
        data: formData,
        contentType: false,
        processData: false,
        success: function(data) {
            listItems()
            $('#form-component').modal('hide');
        },
        error: function(data) {
            showErrorMessages(data)
        }
    });
}

// deleteItem delete an existing item from the API.
const deleteItem = id => {
    $.ajax({
        url: API_URL + "/items/" + id,
        type: 'DELETE',
        contentType: false,
        processData: false,
        success: function(data) {
            listItems()
        },
        error: function(data) {
            console.error(data)
        }
    });
};

// changeOrder patch the order of an existing item in the API.
const changeOrder = (id, order) => {
    formData = new FormData();
    formData.append("order", order);
    $.ajax({
        url: API_URL + "/items/" + id,
        type: 'PATCH',
        data: formData,
        contentType: false,
        processData: false,
        error: function(data) {
            console.error(data)
        }
    });
}

// createSortableList creates a rortable list with jQuery-ui sortable function.
const createSortableList = (e, ui) => {
    $(".sortable").sortable({
        cancel: ".button",
        update: (e, ui) => {
            id = ui.item.data("id");
            order = ui.item.index() + 1;
            changeOrder(id, order);
        }
    });
}

// showErrorMessages handles errors from the API and display them in the form.
const showErrorMessages = (data) => {
    $(".error").html("").hide()
    data.responseJSON.map(error => {
        if (error.field == "image") {
            $(".error.image").html(error.message).show()
        }
        if (error.field == "description") {
            $(".error.description").html(error.message).show()
        }
    })
}

// Initializes onClick and onSubmit handlers, and display all items when docuemnt is ready.
$(function() {
    listItems()

    $(document).on("click", "#new-button", function(e) {
        newItem()
    });

    $(document).on("click", ".edit-button", function(e) {
        editItem($(this).closest('.item').data('id'))
    });

    $(document).on("click", ".delete-button", function(e) {
        deleteItem($(this).closest('.item').data('id'))
    });

    $(document).on("submit", "#new-form", function(e) {
        createItem(e)
    });

    $(document).on("submit", "#edit-form", function(e) {
        updateItem(e, $(this).data('id'))
    });
});