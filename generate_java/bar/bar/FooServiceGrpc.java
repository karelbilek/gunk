package bar;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.35.0)",
    comments = "Source: all.proto")
public final class FooServiceGrpc {

  private FooServiceGrpc() {}

  public static final String SERVICE_NAME = "bar.FooService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      bar.All.Foo> getGetFooMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetFoo",
      requestType = com.google.protobuf.Empty.class,
      responseType = bar.All.Foo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      bar.All.Foo> getGetFooMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, bar.All.Foo> getGetFooMethod;
    if ((getGetFooMethod = FooServiceGrpc.getGetFooMethod) == null) {
      synchronized (FooServiceGrpc.class) {
        if ((getGetFooMethod = FooServiceGrpc.getGetFooMethod) == null) {
          FooServiceGrpc.getGetFooMethod = getGetFooMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, bar.All.Foo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetFoo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  bar.All.Foo.getDefaultInstance()))
              .setSchemaDescriptor(new FooServiceMethodDescriptorSupplier("GetFoo"))
              .build();
        }
      }
    }
    return getGetFooMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FooServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FooServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FooServiceStub>() {
        @java.lang.Override
        public FooServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FooServiceStub(channel, callOptions);
        }
      };
    return FooServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FooServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FooServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FooServiceBlockingStub>() {
        @java.lang.Override
        public FooServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FooServiceBlockingStub(channel, callOptions);
        }
      };
    return FooServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FooServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FooServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FooServiceFutureStub>() {
        @java.lang.Override
        public FooServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FooServiceFutureStub(channel, callOptions);
        }
      };
    return FooServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class FooServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * GetFoo gets foo
     * </pre>
     */
    public void getFoo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<bar.All.Foo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetFooMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetFooMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                bar.All.Foo>(
                  this, METHODID_GET_FOO)))
          .build();
    }
  }

  /**
   */
  public static final class FooServiceStub extends io.grpc.stub.AbstractAsyncStub<FooServiceStub> {
    private FooServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FooServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FooServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetFoo gets foo
     * </pre>
     */
    public void getFoo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<bar.All.Foo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetFooMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class FooServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<FooServiceBlockingStub> {
    private FooServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FooServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FooServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetFoo gets foo
     * </pre>
     */
    public bar.All.Foo getFoo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetFooMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class FooServiceFutureStub extends io.grpc.stub.AbstractFutureStub<FooServiceFutureStub> {
    private FooServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FooServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FooServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetFoo gets foo
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<bar.All.Foo> getFoo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetFooMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_FOO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FooServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FooServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_FOO:
          serviceImpl.getFoo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<bar.All.Foo>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class FooServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    FooServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return bar.All.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("FooService");
    }
  }

  private static final class FooServiceFileDescriptorSupplier
      extends FooServiceBaseDescriptorSupplier {
    FooServiceFileDescriptorSupplier() {}
  }

  private static final class FooServiceMethodDescriptorSupplier
      extends FooServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    FooServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (FooServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FooServiceFileDescriptorSupplier())
              .addMethod(getGetFooMethod())
              .build();
        }
      }
    }
    return result;
  }
}
