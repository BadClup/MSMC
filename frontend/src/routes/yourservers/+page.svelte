<script>
    import Top from "./../Top.svelte";
    import Footer from "./../Footer.svelte";
    import ChooseServer from "./ChooseServer.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
    }
</style>

<div id="container">
    <Top logout={true}/>
    <ChooseServer/>
    <Footer/>
</div>