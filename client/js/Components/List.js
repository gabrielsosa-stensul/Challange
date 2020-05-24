var list = new Component("#list-component", {
    data: {
        items: []
    },
    template: function(props) {
        _items = props.items.map(_item => {
            return (`
                <div class="ui grid item" data-id="` + _item.id + `">
                    <div class="one wide column">
                        <img class="ui avatar image" src="` + API_URL + _item.image + `">
                    </div>
                    <div class="eleven wide column description">
                        ` + _item.description.replace(/\n/g, "<br />") + `
                    </div>
                    <div class="two wide column right floated ui icon buttons">
                        <button class="ui button" onClick="editItem(` + _item.id + `)" title="Edit">
                            <i class="edit icon"></i>
                        </button>
                        <button class="ui button" onClick="deleteItem(` + _item.id + `)">
                            <i class="trash icon"></i>
                        </button>
                    </div>
                </div>
            `);
        });
        return (`
            <div class="ui grid">
                <div class="eight wide column">
                    <div class="ui label large blue">
                        Quantity: <div class="detail">` + props.items.length + `</div>
                    </div>
                </div>
                <div class="eight wide column">
                    <button class="ui right floated small button positive" onclick="return newItem()">
                        New Item
                    </button>
                </div>
            </div>
            <div class="ui big aligned selection divided list sortable">
                ` + _items.join("") + `
            </div>
        `);
    }
});