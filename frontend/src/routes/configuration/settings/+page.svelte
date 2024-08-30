<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import Switch from "$lib/components/Switch.svelte";
    import OptionsList from "./../optionsList.svelte";
</script>
<style>
    section{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    section > div{
        width: 100%;
        flex: 1;
    }

    #listContainer{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    #settings{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 240px);
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
        height: 56px;
        margin-left: 50px;
        margin-bottom: 20px;
        background-color: rgb(48, 48, 48);
    }
    
    .textContainer{
        display: inline-block;
        float: left;
        padding: 13px;
        font-size: 24px;
        color: white;
    }

    .switchContainer{
        margin: 11px;
        float: right;
    }

    .input{
        box-sizing: border-box;
        width: 70px;
        height: 28px;
        float: right;
        margin-top: 14px;
        margin-right: 13px;
        padding-left: 6px;
        font-size: 16px;
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <div id="listContainer">
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </div>
        <div id="settings">
            <h1>Settings</h1>
            <div class="option">
                <div class="textContainer">Option 1</div>
                <div class="switchContainer"><Switch/></div>
            </div>
            <div class="option">
                <div class="textContainer">Option 2</div>
                <input class="input" name="inpu1" type="number" step="1" min="1" max="10" value="1">
            </div>
        </div>
    </div>
    <Footer/>
</section>