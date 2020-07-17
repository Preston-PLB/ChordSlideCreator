<template>
    <div class="settings">
    <form>
        <label class="option">
            <p>Height</p>
            <input type="number" name="height" placeholder="1920">
        </label>
        <label class="option">
            <p>Width</p>
            <input type="number" name="width" placeholder="1080">
        </label>
        <label class="option">
            <p>Nashville Numbers</p>
            <input type="checkbox" name="nashville">
        </label>

        <button id="goButton" onclick="render()">Render</button>
    </form>
    </div>
</template>

<script>
    export default {
        name: "Settings"
    }
</script>

<style scoped>
    .settings{
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

    .option p {
        flex: 2;
        padding: 0 10px 0 0;
        margin: 0
    }

    .option input{
        flex: 1;
        margin: 0;
    }
</style>