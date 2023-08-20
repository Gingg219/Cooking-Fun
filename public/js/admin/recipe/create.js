// Quill Handler

    var toolbarOptions = [
    ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
    ['blockquote', 'code-block'],

    [{ 'header': 1 }, { 'header': 2 }],               // custom button values
    [{ 'list': 'ordered'}, { 'list': 'bullet' }],
    [{ 'script': 'sub'}, { 'script': 'super' }],      // superscript/subscript
    [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent
    [{ 'direction': 'rtl' }],                         // text direction

    [{ 'size': ['small', false, 'large', 'huge'] }],  // custom dropdown
    [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

    [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
    [{ 'font': [] }],
    [{ 'align': [] }],
    ['image'],                                        // add image buttons

    ['clean']                                         // remove formatting button
    ];

    var quill = new Quill('#editor', {
        modules: {
            toolbar: {
                container: toolbarOptions,
            }
        },
        placeholder: 'Enter content here...',
        theme: 'snow'
    });
    quill.on('text-change', function() {
            let value = quill.root.innerHTML
            $('#recipeContent').text(value);
            console.log($('#recipeContent').val());
        });
    $(document).ready(function () {
        $('.ingredient').select2({
            tags: true,
            tokenSeparators : [ ',' , ' ' ],
            width: 'resolve'
        });
    });
    $('#formRecipe').submit(function (e) { 
        e.preventDefault();
        console.log($('.ingredient').val());
        var data = $('#formRecipe').serialize();
        console.log(data);
        $.ajax({
            type: "POST",
            url: "admin/recipe/store",
            data: data,
            dataType: "json",
            success: function (response) {
                alert("Insert successful")
            }
        });
    });
