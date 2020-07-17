<template>
    <div class="files">
    <label class="option vertical-flex">
        <p>Select your files</p>
        <input type="text" name="files" placeholder="~/Documents/Songs/overcome.cho"/>
    </label>
    </div>
</template>

<script>
    export default {
        name: "FilePath"
    }
</script>

<style scoped>
    .files{
        flex: 1;
        padding: 5px;
    }
    .option{
        display: flex;
        align-items: center;
        width: 100%;
        margin: 0;
        padding: 0;
    }
    .vertical-flex{
        flex-direction: column;
    }


    .option p {
        flex: 2;
        padding: 0 10px 0 0;
        margin: 0
    }

    .option input{
        flex: 1;
        margin: 0;
    }

    .vertical-flex{
        flex-direction: column;
    }
</style>