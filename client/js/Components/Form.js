// Form component renders a form to update or create items.
// Its state can contain an item, and in that case it shows a form to update it, otherwise it shows a form to create a new one.
var form = new Component("#form-component", {
    data: {
        item: null
    },
    template: function(props) {
        if (props.item) {
            return `
                <form class="ui form content" id="edit-form" enctype="multipart/form-data" data-id="` + props.item.id + `">
                    <div class="fields">
                        <div class="fifteen wide field">
                            <label>Image</label>
                            <p>jpg, gif, png</p>
                            <input type="file" name="image" accept="image/jpeg,image/gif,image/png">
                            <div class="ui error message image"></div>
                        </div>
                        <div class="one wide field hide">
                            <img src="` + API_URL + props.item.image + `">
                        </div>
                    </div>
                    <div class="field">
                        <label>Description</label>
                        <p>Max 300 chars.</p>
                        <textarea type="text" name="description" maxlength="300" rows="3" required>` + props.item.description + `</textarea>
                        <div class="ui error message description"></div>
                    </div>
                    <button class="ui button positive" type="submit">Update</button>
                </form>
            `;
        } else {
            return `
                <form class="ui form content" id="new-form" enctype="multipart/form-data">
                    <div class="fields">
                        <div class="sixteen wide field">
                            <label>Image</label>
                            <p>jpg, gif, png</p>
                            <input type="file" name="image" accept="image/jpeg,image/gif,image/png" required>
                            <div class="ui error message image"></div>
                        </div>
                    </div>
                    <div class="field">
                        <label>Description</label>
                        <p>Max 300 chars.</p>
                        <textarea type="text" name="description" maxlength="300" rows="3" required></textarea>
                        <div class="ui error message description"></div>
                    </div>
                    <button class="ui button positive" type="submit">Create</button>
                </form>
            `;
        }
    }
});