var form = new Component("#form-component", {
    data: {
        item: null
    },
    template: function(props) {
        if (props.item) {
            return `
                <form class="ui form content" onsubmit="return updateItem(event, ` + props.item.id + `)" enctype="multipart/form-data">
                    <div class="fields">
                        <div class="fifteen wide field">
                            <label>Image</label>
                            <input type="file" name="image" accept="image/jpeg,image/gif,image/png">
                            <div class="ui error message image"></div>
                        </div>
                        <div class="one wide field hide">
                            <img style="height:43px; margin-top: 23px;" src="` + API_URL + props.item.image + `">
                        </div>
                    </div>
                    <div class="field">
                        <label>Description</label>
                        <textarea type="text" name="description" maxlength="300" rows="3" required>` + props.item.description + `</textarea>
                        <div class="ui error message description"></div>
                    </div>
                    <button class="ui button" type="submit">Update</button>
                </form>
            `;
        } else {
            return `
                <form class="ui form content" onsubmit="return createItem(event)" enctype="multipart/form-data">
                    <div class="fields">
                        <div class="sixteen wide field">
                            <label>Image</label>
                            <input type="file" name="image" accept="image/jpeg,image/gif,image/png" required>
                            <div class="ui error message image"></div>
                        </div>
                    </div>
                    <div class="field">
                        <label>Description</label>
                        <textarea type="text" name="description" maxlength="300" rows="3" required></textarea>
                        <div class="ui error message description"></div>
                    </div>
                    <button class="ui button" type="submit">Create</button>
                </form>
            `;
        }
    }
});