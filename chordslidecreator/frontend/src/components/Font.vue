<template>
    <div class="fonts">
        <select>
            <option>Ubuntu</option>
        </select>

        <label class="option">
            <p>Text Color:</p>
            <input type="color" v-model="color" />
            <p>{{ color }}</p>
        </label>

    </div>
</template>

<script>
    export default {
        name: "Font",

        props: {color: String},

    }
</script>

<style scoped>
    .fonts select{
        width: 100%;
    }

    .fonts{
        flex: 1;
        padding: 5px;
        width: 100%;
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
        color: white;
        flex: 1;
        margin: 0;
    }
</style>