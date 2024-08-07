<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import Switch from "$lib/components/Switch.svelte";
    import OptionsList from "./../optionsList.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    #container2{
        width: 100%;
        flex: 1;
    }

    #listContainer{
        display: inline-block;
        float: left;
        height: 100%;
        width: 220px;
    }

    #settings{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 220px);
        float: right;
        padding-top: 15px;
        flex: 1;
    }

    h1{
        position: relative;
        text-align: center;
        font-weight: bold;
        font-size: 50px;
        color: inherit;
    }

    .option{
        box-sizing: border-box;
        width: 320px;
        height: 50px;
        margin-left: 50px;
        background-color: rgb(48, 48, 48);
        color: white;
        font-size: 22px;
    }
    
    .switchContainer{
        padding: 8px;
        float: right;
    }

    .textContainer{
        padding: 10px;
        float: left;
    }
</style>

<div id="container">
    <Top logout={true}/>
    <div id="container2">
        <div id="listContainer">
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </div>
        <div id="settings">
            <h1>Settings</h1><br>
            <div class="option">
                <div class="textContainer">Option 1</div>
                <div class="switchContainer"><Switch/></div>
            </div>
        </div>
    </div>
    <Footer/>
</div>