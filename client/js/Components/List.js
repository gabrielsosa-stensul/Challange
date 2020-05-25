// List component renders an item list, with its respective counter.
// Its state can contain a list of items that will be rendered in an sortable list with its respective button to edit and delete it.
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
                    <div class="twelve wide column description">
                        ` + _item.description.replace(/\n/g, "<br />") + `
                    </div>
                    <div class="two wide column right floated ui icon buttons">
                        <button class="ui button edit-button">
                            <i class="edit icon"></i>
                        </button>
                        <button class="ui button delete-button">
                            <i class="trash icon"></i>
                        </button>
                    </div>
                </div>
            `);
        });
        return (`
            <div class="ui label large blue">
                Quantity: <div class="detail">` + props.items.length + `</div>
            </div>
            <div class="ui big aligned selection divided list sortable">
                ` + _items.join("") + `
            </div>
        `);
    }
});