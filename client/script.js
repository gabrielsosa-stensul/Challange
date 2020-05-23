const API_URL = "/api"

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

const newItem = () => {
    form.data.item = null;
    form.render();
    $('#form-component').modal('show');
}

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

const createSortableList = (event, ui) => {
    $(".sortable").sortable({
        cancel: ".button",
        update: (event, ui) => {
            order = ui.item.index() + 1;
            id = ui.item.data("id");
            changeOrder(order, id);
        }
    });
}

const showErrorMessages = (data) => {
    if (data.responseJSON.image) {
        $(".error.image").html(data.responseJSON.image).show()
    }
    if (data.responseJSON.description) {
        $(".error.description").html(data.responseJSON.description).show()
    }
}

$(function() {
    listItems()
});