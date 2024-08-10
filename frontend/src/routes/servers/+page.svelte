<script>
    import Top from "$lib/components//Top.svelte";
    import Footer from "$lib/components//Footer.svelte";
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
    <Top canLogout={true}/>
    <ChooseServer/>
    <Footer/>
</div>