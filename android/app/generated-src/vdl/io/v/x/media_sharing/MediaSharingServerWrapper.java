// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.

// Source(s):  media.vdl
package io.v.x.media_sharing;

/**
 * Wrapper for {@link MediaSharingServer}.  This wrapper is used by
 * {@link io.v.v23.rpc.ReflectInvoker} to indirectly invoke server methods.
 */
public final class MediaSharingServerWrapper {
    private final io.v.x.media_sharing.MediaSharingServer server;




    /**
     * Creates a new {@link MediaSharingServerWrapper} to invoke the methods of the
     * provided server.
     *
     * @param server server whose methods are to be invoked
     */
    public MediaSharingServerWrapper(io.v.x.media_sharing.MediaSharingServer server) {
        this.server = server;
        
        
    }

    /**
     * Returns a description of this server.
     */
    public io.v.v23.vdlroot.signature.Interface signature() {
        java.util.List<io.v.v23.vdlroot.signature.Embed> embeds = new java.util.ArrayList<io.v.v23.vdlroot.signature.Embed>();
        java.util.List<io.v.v23.vdlroot.signature.Method> methods = new java.util.ArrayList<io.v.v23.vdlroot.signature.Method>();
        
        {
            java.util.List<io.v.v23.vdlroot.signature.Arg> inArgs = new java.util.ArrayList<io.v.v23.vdlroot.signature.Arg>();
            
            inArgs.add(new io.v.v23.vdlroot.signature.Arg("", "", new io.v.v23.vdl.VdlTypeObject(new com.google.common.reflect.TypeToken<java.lang.String>(){}.getType())));
            
            java.util.List<io.v.v23.vdlroot.signature.Arg> outArgs = new java.util.ArrayList<io.v.v23.vdlroot.signature.Arg>();
            
            java.util.List<io.v.v23.vdl.VdlAny> tags = new java.util.ArrayList<io.v.v23.vdl.VdlAny>();
            
            methods.add(new io.v.v23.vdlroot.signature.Method(
                "displayUrl",
                "// DisplayURL will cause the server to display whatever media is at" + 
"// the given URL.  The server will rely on the ContentType response" + 
"// header it gets when fetching the url to decide how to display" + 
"// the media." + 
"",
                inArgs,
                outArgs,
                null,
                null,
                tags));
        }
        
        {
            java.util.List<io.v.v23.vdlroot.signature.Arg> inArgs = new java.util.ArrayList<io.v.v23.vdlroot.signature.Arg>();
            
            inArgs.add(new io.v.v23.vdlroot.signature.Arg("", "", new io.v.v23.vdl.VdlTypeObject(new com.google.common.reflect.TypeToken<java.lang.String>(){}.getType())));
            
            java.util.List<io.v.v23.vdlroot.signature.Arg> outArgs = new java.util.ArrayList<io.v.v23.vdlroot.signature.Arg>();
            
            java.util.List<io.v.v23.vdl.VdlAny> tags = new java.util.ArrayList<io.v.v23.vdl.VdlAny>();
            
            methods.add(new io.v.v23.vdlroot.signature.Method(
                "displayBytes",
                "// DisplayBytes will cause the server to display whatever media is" + 
"// sent in the stream.  In the case of audio or movie media, the" + 
"// media should be played while the data is streaming.  The mediaType" + 
"// can be used by the server to decide how to display the media." + 
"",
                inArgs,
                outArgs,
                null,
                null,
                tags));
        }
        

        return new io.v.v23.vdlroot.signature.Interface("MediaSharing", "io.v.x.media_sharing", "", embeds, methods);
    }

    /**
     * Returns all tags associated with the provided method or {@code null} if the method isn't
     * implemented by this server.
     *
     * @param method method whose tags are to be returned
     */
    @SuppressWarnings("unused")
    public io.v.v23.vdl.VdlValue[] getMethodTags(java.lang.String method) throws io.v.v23.verror.VException {
        
        if ("displayBytes".equals(method)) {
            try {
                return new io.v.v23.vdl.VdlValue[] {
                    
                };
            } catch (IllegalArgumentException e) {
                throw new io.v.v23.verror.VException(String.format("Couldn't get tags for method \"displayBytes\": %s", e.getMessage()));
            }
        }
        
        if ("displayUrl".equals(method)) {
            try {
                return new io.v.v23.vdl.VdlValue[] {
                    
                };
            } catch (IllegalArgumentException e) {
                throw new io.v.v23.verror.VException(String.format("Couldn't get tags for method \"displayUrl\": %s", e.getMessage()));
            }
        }
        
        
        return null;  // method not found
    }

     
    
    /**
 * DisplayURL will cause the server to display whatever media is at
 * the given URL.  The server will rely on the ContentType response
 * header it gets when fetching the url to decide how to display
 * the media.
*/
    public void displayUrl(io.v.v23.context.VContext ctx, final io.v.v23.rpc.StreamServerCall call, final java.lang.String url) throws io.v.v23.verror.VException {
         
         this.server.displayUrl(ctx, call , url  );
    }

    /**
 * DisplayBytes will cause the server to display whatever media is
 * sent in the stream.  In the case of audio or movie media, the
 * media should be played while the data is streaming.  The mediaType
 * can be used by the server to decide how to display the media.
*/
    public void displayBytes(io.v.v23.context.VContext ctx, final io.v.v23.rpc.StreamServerCall call, final java.lang.String mediaType) throws io.v.v23.verror.VException {
        
        io.v.v23.vdl.Stream<java.lang.Void, byte[]> _stream = new io.v.v23.vdl.Stream<java.lang.Void, byte[]>() {
            @Override
            public void send(java.lang.Void item) throws io.v.v23.verror.VException {
                java.lang.reflect.Type type = new com.google.common.reflect.TypeToken< java.lang.Void >() {}.getType();
                call.send(item, type);
            }
            @Override
            public byte[] recv() throws java.io.EOFException, io.v.v23.verror.VException {
                java.lang.reflect.Type type = new com.google.common.reflect.TypeToken< byte[] >() {}.getType();
                java.lang.Object result = call.recv(type);
                try {
                    return (byte[])result;
                } catch (java.lang.ClassCastException e) {
                    throw new io.v.v23.verror.VException("Unexpected result type: " + result.getClass().getCanonicalName());
                }
            }
        };
         
         this.server.displayBytes(ctx, call , mediaType  ,_stream  );
    }



 

}