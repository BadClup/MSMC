<script lang="ts">
    export let state = false;
</script>
<style>
    .switch {
        position: relative;
        display: inline-block;
        width: 60px;
        height: 34px;
    }

    .switch input {
        display: inline-block;
        opacity: 0;
        width: 0;
        height: 0;
    }

    .slider {
        display: inline-block;
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: red;
        transition: 0.5s;
        border-radius: 34px;
    }

    .slider:before {
        display: inline-block;
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: 0.5s;
        border-radius: 50%;
    }

    input:checked + .slider {
        background-color: #00e600;
    }

    input:checked + .slider:before {
        transform: translateX(26px);
    }
</style>


<div>
    <label class="switch">
        {#if state == true}
            <input type="checkbox" checked>
        {:else}
            <input type="checkbox">
        {/if}
        <span class="slider"></span>
    </label>
</div>